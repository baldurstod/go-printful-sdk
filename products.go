package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) GetCatalogProduct(productId int, opts ...RequestOption) (*model.Product, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	headers := map[string]string{}
	if opt.language != "" {
		headers["X-PF-Language"] = opt.language
	}

	u, _ := buildURL("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(productId), opt)
	log.Println(u)
	resp, err := c.Get(u, headers, ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to get printful response")
	}
	defer resp.Body.Close()

	response := &responses.ProductResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	return &response.Data.Product, nil
}

func (c *PrintfulClient) GetProductCategories(productId int, opts ...RequestOption) ([]model.Category, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	headers := map[string]string{}
	if opt.language != "" {
		headers["X-PF-Language"] = opt.language
	}

	categories := make([]model.Category, 0, 100)
	for {
		u, _ := buildURL("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(productId)+"/catalog-categories", opt)
		resp, err := c.Get(u, headers, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.CategoriesResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		categories = append(categories, response.Data...)

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return categories, nil
}
