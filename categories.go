package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) GetCatalogCategories(opts ...requestOption) ([]model.Category, error) {
	opt := getOptions(opts...)

	categories := make([]model.Category, 0, 400)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100

	for {
		u, _ := buildURL("https://api.printful.com/v2/catalog-categories", opt)
		resp, err := c.Get(u, nil, ctx)
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
