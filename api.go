package printfulsdk

type RequestBodyKey int64

const (
	FileRole RequestBodyKey = iota
	URL
	Filename
	FileVisible
	OrderExternalID
	OrderShippingMethod
	OrderCustomization
	OrderRetailCosts
)

func BuildRequestBody(o options, keys ...RequestBodyKey) map[string]interface{} {
	body := map[string]interface{}{}

	for _, key := range keys {
		switch key {
		case FileRole:
			if o.fileRole != "" {
				body["role"] = o.fileRole
			}
		case URL:
			body["url"] = o.url
		case Filename:
			if o.filename != "" {
				body["filename"] = o.filename
			}
		case FileVisible:
			body["visible"] = o.fileVisible
		case OrderExternalID:
			if o.orderExternalID != "" {
				body["external_id"] = o.orderExternalID
			}
		case OrderShippingMethod:
			if o.orderShippingMethod != "" {
				body["shipping"] = o.orderShippingMethod
			}
		case OrderCustomization:
			if o.orderCustomization != nil {
				body["customization"] = o.orderCustomization
			}
		case OrderRetailCosts:
			if o.orderRetailCosts != nil {
				body["retail_costs"] = o.orderRetailCosts
			}
		}
	}
	return body
}
