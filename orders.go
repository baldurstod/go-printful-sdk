package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) CreateOrder(recipient model.Address, items []model.CatalogItem, opts ...RequestOption) (*model.Order, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	body := BuildRequestBody(opt, OrderExternalID, OrderShippingMethod, OrderCustomization, OrderRetailCosts)

	body["recipient"] = recipient
	body["order_items"] = items

	//b, _ := json.MarshalIndent(body, "", "  ")
	//log.Println(string(b))

	u := "https://api.printful.com/v2/orders"
	resp, err := c.Post(u, nil, body, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("post returned an error in CreateOrder: %w", err)
	}

	response := &responses.CreateOrderResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	return &response.Data, nil
}

func (c *PrintfulClient) GetOrder(id int, opts ...RequestOption) (*model.Order, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	u, _ := buildURL("https://api.printful.com/v2/orders/"+strconv.Itoa(id), opt)
	resp, err := c.Get(u, nil, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("get returned an error in GetOrder: %w", err)
	}

	response := &responses.CreateOrderResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	return &response.Data, nil
}

func (c *PrintfulClient) GetOrderItem(orderID int, itemID int, opts ...RequestOption) (*model.CatalogItemReadonly, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	u, _ := buildURL("https://api.printful.com/v2/orders/"+strconv.Itoa(orderID)+"/order-items/"+strconv.Itoa(itemID), opt)
	resp, err := c.Get(u, nil, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("get returned an error in GetOrderItem: %w", err)
	}

	response := &responses.GetItemById{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	return &response.Data, nil
}
