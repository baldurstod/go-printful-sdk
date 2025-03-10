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

type requestOption func(*options)

func GetOptions(opts ...requestOption) options {
	return getOptions(opts...)
}

func getOptions(opts ...requestOption) options {
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

func WithOffset(offset uint) requestOption {
	return func(o *options) {
		o.offset = offset
	}
}

func WithLimit(limit uint) requestOption {
	return func(o *options) {
		o.limit = limit
	}
}

func WithCategories(categories ...int) requestOption {
	return func(o *options) {
		o.categories = append(o.categories, categories...)
	}
}

func WithColors(colors ...string) requestOption {
	return func(o *options) {
		o.colors = append(o.colors, colors...)
	}
}

func WithPlacements(placements ...string) requestOption {
	return func(o *options) {
		o.placements = append(o.placements, placements...)
	}
}

func WithOnlyNew() requestOption {
	return func(o *options) {
		o.new = true
	}
}

func WithSellingRegionName(sellingRegionName string) requestOption {
	return func(o *options) {
		o.sellingRegionName = sellingRegionName
	}
}

func WithCurrency(currency string) requestOption {
	return func(o *options) {
		o.currency = currency
	}
}

func WithSortDirection(sortDirection SortDirection) requestOption {
	return func(o *options) {
		o.sortDirection = sortDirection
	}
}

func WithSortType(sortType SortType) requestOption {
	return func(o *options) {
		o.sortType = sortType
	}
}

func WithTechniques(techniques ...Technique) requestOption {
	return func(o *options) {
		o.techniques = append(o.techniques, techniques...)
	}
}

func WithLanguage(language string) requestOption {
	return func(o *options) {
		o.language = language
	}
}

func WithTimeout(timeout time.Duration) requestOption {
	return func(o *options) {
		o.timeout = timeout
	}
}

func SetFileRole(role string) requestOption {
	return func(o *options) {
		o.fileRole = role
	}
}

func SetURL(url string) requestOption {
	return func(o *options) {
		o.url = url
	}
}

func SetFilename(filename string) requestOption {
	return func(o *options) {
		o.filename = filename
	}
}

func SetFileVisible(visible bool) requestOption {
	return func(o *options) {
		o.fileVisible = visible
	}
}

func SetOrderExternalID(externalID string) requestOption {
	return func(o *options) {
		o.orderExternalID = externalID
	}
}

func SetOrderShippingMethod(shippingMethod string) requestOption {
	return func(o *options) {
		o.orderShippingMethod = shippingMethod
	}
}

func SetOrderCustomization(customization *model.Customization) requestOption {
	return func(o *options) {
		o.orderCustomization = customization
	}
}

func SetOrderRetailCosts(retailCosts *model.RetailCosts2) requestOption {
	return func(o *options) {
		o.orderRetailCosts = retailCosts
	}
}
