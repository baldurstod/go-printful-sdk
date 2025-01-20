package printfulapi

import "time"

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
	DtFilm      SortType  = "dtfilm"
)

type options struct {
	categories    []int
	colors        []string
	placements    []string
	techniques    []Technique
	offset        uint
	limit         uint
	new           bool
	regionName    string
	sortDirection SortDirection
	sortType      SortType
	language      string
	timeout       time.Duration
}

type requestOption func(*options)

func WithOffset(offset uint) requestOption {
	return func(o *options) {
		o.offset = offset
	}
}

func WithLimit(limit uint) requestOption {
	return func(o *options) {
		if limit <= 100 {
			o.limit = limit
		} else {
			o.limit = 100
		}
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
		o.colors = append(o.placements, placements...)
	}
}

func WithOnlyNew() requestOption {
	return func(o *options) {
		o.new = true
	}
}

func WithRegionName(regionName string) requestOption {
	return func(o *options) {
		o.regionName = regionName
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

func getOptions(opts ...requestOption) options {
	cfg := options{
		limit:   20,
		timeout: time.Duration(-1),
	}
	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}
