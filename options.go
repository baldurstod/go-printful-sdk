package printfulsdk

import (
	"time"

	"github.com/baldurstod/go-printful-sdk/model"
)

type SortDirection string

const (
	SortAscending  SortDirection = "ascending"
	SortDescending SortDirection = "descending"
)

type SortType string

const (
	SortNew        SortType = "new"
	SortRating     SortType = "rating"
	SortPrice      SortType = "price"
	SortBestseller SortType = "bestseller"
)

type Technique string

const (
	Dtg         Technique = "dtg"
	Digital     Technique = "digital"
	CutSew      Technique = "cut-sew"
	Uv          Technique = "uv"
	Embroidery  Technique = "embroidery"
	Sublimation Technique = "sublimation"
	DtFilm      Technique = "dtfilm"
)

type options struct {
	categories          []int
	colors              []string
	placements          []string
	techniques          []Technique
	offset              uint
	limit               uint
	new                 bool
	sellingRegionName   string
	currency            string
	sortDirection       SortDirection
	sortType            SortType
	language            string
	timeout             time.Duration
	url                 string
	fileRole            string
	filename            string
	fileVisible         bool
	orderExternalID     string
	orderShippingMethod string
	orderCustomization  *model.Customization
	orderRetailCosts    *model.RetailCosts2
}

type RequestOption func(*options)

func GetOptions(opts ...RequestOption) options {
	return getOptions(opts...)
}

func getOptions(opts ...RequestOption) options {
	cfg := options{
		limit:       0,
		timeout:     time.Duration(-1),
		fileVisible: true,
		currency:    "USD",
	}
	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}

func WithOffset(offset uint) RequestOption {
	return func(o *options) {
		o.offset = offset
	}
}

func WithLimit(limit uint) RequestOption {
	return func(o *options) {
		o.limit = limit
	}
}

func WithCategories(categories ...int) RequestOption {
	return func(o *options) {
		o.categories = append(o.categories, categories...)
	}
}

func WithColors(colors ...string) RequestOption {
	return func(o *options) {
		o.colors = append(o.colors, colors...)
	}
}

func WithPlacements(placements ...string) RequestOption {
	return func(o *options) {
		o.placements = append(o.placements, placements...)
	}
}

func WithOnlyNew() RequestOption {
	return func(o *options) {
		o.new = true
	}
}

func WithSellingRegionName(sellingRegionName string) RequestOption {
	return func(o *options) {
		o.sellingRegionName = sellingRegionName
	}
}

func WithCurrency(currency string) RequestOption {
	return func(o *options) {
		o.currency = currency
	}
}

func WithSortDirection(sortDirection SortDirection) RequestOption {
	return func(o *options) {
		o.sortDirection = sortDirection
	}
}

func WithSortType(sortType SortType) RequestOption {
	return func(o *options) {
		o.sortType = sortType
	}
}

func WithTechniques(techniques ...Technique) RequestOption {
	return func(o *options) {
		o.techniques = append(o.techniques, techniques...)
	}
}

func WithLanguage(language string) RequestOption {
	return func(o *options) {
		o.language = language
	}
}

func WithTimeout(timeout time.Duration) RequestOption {
	return func(o *options) {
		o.timeout = timeout
	}
}

func SetFileRole(role string) RequestOption {
	return func(o *options) {
		o.fileRole = role
	}
}

func SetURL(url string) RequestOption {
	return func(o *options) {
		o.url = url
	}
}

func SetFilename(filename string) RequestOption {
	return func(o *options) {
		o.filename = filename
	}
}

func SetFileVisible(visible bool) RequestOption {
	return func(o *options) {
		o.fileVisible = visible
	}
}

func SetOrderExternalID(externalID string) RequestOption {
	return func(o *options) {
		o.orderExternalID = externalID
	}
}

func SetOrderShippingMethod(shippingMethod string) RequestOption {
	return func(o *options) {
		o.orderShippingMethod = shippingMethod
	}
}

func SetOrderCustomization(customization *model.Customization) RequestOption {
	return func(o *options) {
		o.orderCustomization = customization
	}
}

func SetOrderRetailCosts(retailCosts *model.RetailCosts2) RequestOption {
	return func(o *options) {
		o.orderRetailCosts = retailCosts
	}
}
