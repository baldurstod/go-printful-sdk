package printfulapi

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
	"time"

	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
)

const PRINTFUL_CATALOG_ENDPOINT = "https://api.printful.com/v2/catalog-products"
const PRINTFUL_ORDERS_ENDPOINT = "https://api.printful.com/v2/orders"
const PRINTFUL_FILES_ENDPOINT = "https://api.printful.com/v2/files"
const PRINTFUL_COUNTRIES_ENDPOINT = "https://api.printful.com/v2/countries"
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

func (c *PrintfulClient) get(endpoint string, path string, headers map[string]string, ctx context.Context) (*http.Response, error) {
	return c.fetch("GET", endpoint, path, headers, nil, ctx)
}

func (c *PrintfulClient) post(endpoint string, path string, headers map[string]string, body map[string]interface{}, ctx context.Context) (*http.Response, error) {
	return c.fetch("POST", endpoint, path, headers, body, ctx)
}

func (c *PrintfulClient) fetch(method string, endpoint string, path string, headers map[string]string, body map[string]interface{}, ctx context.Context) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var limiter *rate.Limiter

	if method == "POST" && endpoint == PRINTFUL_MOCKUP_ENDPOINT {
		limiter = c.mockupLimiter
	} else {
		limiter = c.stdLimiter
	}

	u, err := url.JoinPath(endpoint, path)
	if err != nil {
		return nil, errors.New("unable to create URL")
	}

	var requestBody io.Reader
	if body != nil {
		out, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		requestBody = bytes.NewBuffer(out)
	}

	var resp *http.Response
	req, err := http.NewRequestWithContext(ctx, method, u, requestBody)
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

	if resp.StatusCode != 200 { //Everything except 429 and 200
		if resp.StatusCode == 429 { //Too Many Requests
			fmt.Println("429", endpoint, header.Get("X-RateLimit-Remaining"), header.Get("X-RateLimit-Reset"), header.Get("X-RateLimit-Limit"), header.Get("X-RateLimit-Policy"), header.Get("retry-after"))
		}
		return nil, fmt.Errorf("printful returned HTTP status code: %d", resp.StatusCode)
	}

	//fmt.Println("remaining", endpoint, header.Get("X-RateLimit-Remaining"), header.Get("X-RateLimit-Reset"), header.Get("X-RateLimit-Limit"))

	return resp, err
}

func (c *PrintfulClient) GetCatalogProducts(opts ...requestOption) error {
	opt := getOptions(opts...)

	var ctx context.Context
	if opt.timeout > 0 {
		ctx, _ = context.WithTimeout(context.Background(), opt.timeout)
	}

	resp, err := c.get(PRINTFUL_CATALOG_ENDPOINT, "", nil, ctx)
	if err != nil {
		log.Println(err)
		return errors.New("unable to get printful response")
	}

	response := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return errors.New("unable to decode printful response")
	}

	return nil
}

func (c *PrintfulClient) GetCountries(opts ...requestOption) error {
	opt := getOptions(opts...)

	var ctx context.Context
	if opt.timeout > 0 {
		ctx, _ = context.WithTimeout(context.Background(), opt.timeout)
	}

	resp, err := c.get(PRINTFUL_COUNTRIES_ENDPOINT, "", nil, ctx)
	if err != nil {
		log.Println(err)
		return errors.New("unable to get printful response")
	}

	response := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return errors.New("unable to decode printful response")
	}

	return nil
}
