package printfulsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
)

const PRINTFUL_CATALOG_PRODUCTS = "https://api.printful.com/v2/catalog-products"
const PRINTFUL_CATALOG_VARIANTS = "https://api.printful.com/v2/catalog-variants"
const PRINTFUL_ORDERS_ENDPOINT = "https://api.printful.com/v2/orders"
const PRINTFUL_FILES_ENDPOINT = "https://api.printful.com/v2/files"
const PRINTFUL_COUNTRIES = "https://api.printful.com/v2/countries"
const PRINTFUL_SHIPPING_RATES_ENDPOINT = "https://api.printful.com/v2/shipping-rates"
const PRINTFUL_MOCKUP_ENDPOINT = "https://api.printful.com/v2/mockup-tasks"
const PRINTFUL_STORES_ENDPOINT = "https://api.printful.com/v2/stores"
const PRINTFUL_APPROVAL_SHEETS_ENDPOINT = "https://api.printful.com/v2/approval-sheets"

type PrintfulClient struct {
	accessToken   string
	stdLimiter    *rate.Limiter
	mockupLimiter *rate.Limiter
	sem           *semaphore.Weighted
}

func NewPrintfulClient(accessToken string) *PrintfulClient {
	return &PrintfulClient{
		accessToken: accessToken,
		// Notice: these values will be updated depending on returned X-Ratelimit headers
		stdLimiter:    rate.NewLimiter(2, 120),
		mockupLimiter: rate.NewLimiter(1./30., 2),
		sem:           semaphore.NewWeighted(int64(20)),
	}
}

// Change access token. Any queued request still uses the old token
func (c *PrintfulClient) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

func (c *PrintfulClient) Get(path string, headers map[string]string, ctx context.Context) (*http.Response, error) {
	return c.fetch("GET", path, headers, nil, ctx)
}

func (c *PrintfulClient) Post(path string, headers map[string]string, body map[string]interface{}, ctx context.Context) (*http.Response, error) {
	return c.fetch("POST", path, headers, body, ctx)
}

func (c *PrintfulClient) fetch(method string, path string, headers map[string]string, body map[string]interface{}, ctx context.Context) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var limiter *rate.Limiter

	if method == "POST" && strings.HasPrefix(path, PRINTFUL_MOCKUP_ENDPOINT) {
		limiter = c.mockupLimiter
	} else {
		limiter = c.stdLimiter
	}

	//u, err := url.parse(endpoint, path)
	/*
		if err != nil {
			return nil, errors.New("unable to create URL")
		}
	*/

	var requestBody io.Reader
	if body != nil {
		out, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		requestBody = bytes.NewBuffer(out)
	}

	var resp *http.Response
	req, err := http.NewRequestWithContext(ctx, method, path, requestBody)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Adding OAuth token
	req.Header.Add("Authorization", "Bearer "+c.accessToken)

	var header http.Header
	for i := 0; i < 10; i++ {
		// Wait for a rate limit token
		err = limiter.Wait(ctx)
		if err != nil {
			return nil, err
		}

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		header = resp.Header

		// Check remaining tokens
		if remaining, err := strconv.Atoi(header.Get("X-RateLimit-Remaining")); err == nil {
			tokens := int(limiter.Tokens())
			if tokens > remaining {
				// Synchronize limiter
				limiter.ReserveN(time.Now(), tokens-remaining)
			}
		}

		r := getRateFromPolicy(header.Get("X-RateLimit-Policy"))
		if r > 0 {
			limiter.SetLimit(rate.Limit(r))
		}

		if resp.StatusCode != 429 {
			// Exit the loop unless we have a code 429 Too Many Requests
			break
		}
	}

	if resp.StatusCode != 200 {
		switch {
		case resp.StatusCode == 429:
			//log.Println("429", path, header.Get("X-RateLimit-Remaining"), header.Get("X-RateLimit-Reset"), header.Get("X-RateLimit-Limit"), header.Get("X-RateLimit-Policy"), header.Get("retry-after"))
			response := map[string]interface{}{
				path:          path,
				"remaining":   header.Get("X-RateLimit-Remaining"),
				"reset":       header.Get("X-RateLimit-Reset"),
				"limit":       header.Get("X-RateLimit-Limit"),
				"policy":      header.Get("X-RateLimit-Policy"),
				"retry-after": header.Get("retry-after"),
			}
			return nil, NewHTTPError(fmt.Errorf("printful returned HTTP status code: %d", resp.StatusCode), response)
		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			response := &responses.Error4XXResponse{}
			if err = json.NewDecoder(resp.Body).Decode(&response); err == nil {
				return nil, NewHTTPError(fmt.Errorf("printful returned HTTP status code: %d", resp.StatusCode), response)
			}
		case resp.StatusCode >= 500 && resp.StatusCode < 600:
			response := &responses.Error5XXResponse{}
			if err = json.NewDecoder(resp.Body).Decode(&response); err == nil {
				return nil, NewHTTPError(fmt.Errorf("printful returned HTTP status code: %d", resp.StatusCode), response)
			}
		}
		return nil, fmt.Errorf("printful returned HTTP status code: %d", resp.StatusCode)
	}

	//log.Println("remaining", path, header.Get("X-RateLimit-Remaining"), header.Get("X-RateLimit-Reset"), header.Get("X-RateLimit-Limit"))

	return resp, err
}

func buildURL(path string, o options) (string, error) {
	u, err := url.ParseRequestURI(path)
	q := url.Values{}
	if err != nil {
		return "", err
	}

	if o.limit != 0 {
		q.Set("limit", strconv.Itoa(int(o.limit)))
	}

	if o.offset != 0 {
		q.Set("offset", strconv.Itoa(int(o.offset)))
	}

	if o.sellingRegionName != "" {
		q.Set("selling_region_name", o.sellingRegionName)
	}

	if o.currency != "" {
		q.Set("currency", o.currency)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (c *PrintfulClient) GetCatalogProducts(opts ...requestOption) ([]model.Product, error) {
	opt := getOptions(opts...)

	products := make([]model.Product, 0, 400)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {

		u, _ := buildURL(PRINTFUL_CATALOG_PRODUCTS, opt)
		log.Println(u)
		resp, err := c.Get(u, nil, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.ProductsResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		products = append(products, response.Data...)

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return products, nil
}

func (c *PrintfulClient) GetCatalogVariants(productId int, opts ...requestOption) ([]model.Variant, error) {
	opt := getOptions(opts...)

	variants := make([]model.Variant, 0, 10)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {
		u, _ := buildURL("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(productId)+"/catalog-variants", opt)
		resp, err := c.Get(u, nil, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.VariantsResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		variants = append(variants, response.Data...)

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return variants, nil
}

func (c *PrintfulClient) GetProductPrices(productId int, opts ...requestOption) (*model.ProductPrices, error) {
	opt := getOptions(opts...)

	prices := model.ProductPrices{}

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {
		u, _ := buildURL("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(productId)+"/prices", opt)
		resp, err := c.Get(u, nil, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.ProductPricesResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		prices.Currency = response.Data.Currency
		prices.Product = response.Data.Product
		prices.Variants = append(prices.Variants, response.Data.Variants...)

		next := response.Paging.Offset + response.Paging.Limit

		if next >= response.Paging.Total {
			break
		}

		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return &prices, nil
}

func (c *PrintfulClient) GetVariantPrices(varianttId int, opts ...requestOption) (*model.VariantPrice, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	u, _ := buildURL("https://api.printful.com/v2/catalog-variants/"+strconv.Itoa(varianttId)+"/prices", opt)
	resp, err := c.Get(u, nil, ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to get printful response")
	}
	defer resp.Body.Close()

	response := &responses.VariantPricesResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	//variants = append(variants, response.Data...)

	return &response.Data, nil
}

func (c *PrintfulClient) GetCountries(opts ...requestOption) ([]model.Country, error) {
	opt := getOptions(opts...)

	countries := make([]model.Country, 0, 200)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {
		u, _ := buildURL(PRINTFUL_COUNTRIES, opt)
		log.Println(u)
		resp, err := c.Get(u, nil, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.CountriesResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		countries = append(countries, response.Data...)

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return countries, nil
}

func (c *PrintfulClient) GetMockupTemplates(productId int, opts ...requestOption) ([]model.MockupTemplates, error) {
	opt := getOptions(opts...)

	templates := make([]model.MockupTemplates, 0, 10)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {
		u, _ := buildURL("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(productId)+"/mockup-templates", opt)
		log.Println(u)
		resp, err := c.Get(u, nil, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.MockupTemplatesResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		templates = append(templates, response.Data...)

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return templates, nil
}

func (c *PrintfulClient) GetMockupStyles(productId int, opts ...requestOption) ([]model.MockupStyles, error) {
	opt := getOptions(opts...)

	styles := make([]model.MockupStyles, 0, 10)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {
		u, _ := buildURL("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(productId)+"/mockup-styles", opt)
		log.Println(u)
		resp, err := c.Get(u, nil, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.MockupStylesResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		styles = append(styles, response.Data...)

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return styles, nil
}
