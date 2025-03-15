package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
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

// Returns order/item id or external id depending on the parameter type
func getId(id any) (string, error) {
	switch id := id.(type) {
	case int:
		return strconv.Itoa(id), nil
	case string:
		return "@" + url.PathEscape(id), nil
	default:
		return "", errors.New("order type must be int or string")
	}
}

// GetOrder return the printful order by id or external id
// if orderID is an integer, returns the order by printful id
// if orderID is a string, returns the order by external id
func (c *PrintfulClient) GetOrder(orderID any, opts ...RequestOption) (*model.Order, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	id, err := getId(orderID)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error while formatting order id in GetOrder: %w", err)
	}

	u, _ := buildURL("https://api.printful.com/v2/orders/"+id, opt)
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

func (c *PrintfulClient) GetOrderItem(orderID any, itemID any, opts ...RequestOption) (*model.CatalogItemReadonly, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	id, err := getId(orderID)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error while formatting order id in GetOrderItem: %w", err)
	}

	id2, err := getId(itemID)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error while formatting item id in GetOrderItem: %w", err)
	}

	u, _ := buildURL("https://api.printful.com/v2/orders/"+id+"/order-items/"+id2, opt)
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
