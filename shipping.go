package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) CalculateShippingRates(recipient model.ShippingRatesAddress, items []model.CatalogOrWarehouseShippingRateItem, opts ...RequestOption) ([]model.ShippingRate, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	body := BuildRequestBody(opt, OrderCurrency)

	body["recipient"] = recipient
	body["order_items"] = items

	u := "https://api.printful.com/v2/shipping-rates"
	resp, err := c.Post(u, nil, body, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("post returned an error in CalculateShippingRates: %w", err)
	}

	response := &responses.ShippingRatesResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	return response.Data, nil
}

/*



	body := map[string]interface{}{}
	err := mapstructure.Decode(datas, &body)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error while decoding params")
	}

	log.Println(body)

	headers := map[string]string{
		"Authorization": "Bearer " + printfulConfig.AccessToken,
	}

	resp, err := fetchRateLimited("POST", PRINTFUL_SHIPPING_API, "/shipping-rates", headers, body)
	if err != nil {
		return nil, errors.New("unable to get printful response")
	}
	defer resp.Body.Close()

	//response := map[string]interface{}{}
	response := responses.ShippingRates{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}
	log.Println(response)

	//p := &(response.Result)

	return response.Result, nil
}
*/
