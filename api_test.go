package printfulsdk_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"image/png"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	printfulsdk "github.com/baldurstod/go-printful-sdk"
	"github.com/baldurstod/go-printful-sdk/model"
)

type Config struct {
	Printful `json:"printful"`
}

type Printful struct {
	AccessToken string `json:"access_token"`
}

var client *printfulsdk.PrintfulClient

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		panic(err)
	}

	client = printfulsdk.NewPrintfulClient(token)
}

func getAuthToken() (string, error) {
	config := Config{}
	var err error
	if content, err := os.ReadFile("config.json"); err == nil {
		if err = json.Unmarshal(content, &config); err == nil {
			return config.AccessToken, nil
		}
	}
	return "", err
}

func TestRateLimiter(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	client.GetCountries()
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup

	var done = 0
	for i := 1; i <= 130; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			client.GetCountries( /*printfulsdk.WithTimeout(1 * time.Second)*/ )
			done = done + 1
			//log.Println(done)
		}()
	}

	wg.Wait()

	_, err = client.GetCatalogProducts()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetProducts(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	products, err := client.GetCatalogProducts( /*printfulsdk.WithLimit(100)*/ /*, printfulsdk.WithTimeout(5*time.Second)*/ )
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&products, "", "\t")

	err = os.WriteFile("./var/products.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetProduct(t *testing.T) {
	id := 785
	product, err := client.GetCatalogProduct(id)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&product, "", "\t")

	err = os.WriteFile("./var/product_"+strconv.Itoa(id)+".json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetVariants(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	products, err := client.GetCatalogVariants(71)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&products, "", "\t")

	err = os.WriteFile("./var/variants.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetProductPrices(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	products, err := client.GetProductPrices(71, printfulsdk.WithCurrency("EUR") /*, printfulsdk.WithSellingRegionName("new_zealand")*/)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&products, "", "\t")

	err = os.WriteFile("./var/product_prices.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetVariantPrices(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	products, err := client.GetVariantPrices(17008, printfulsdk.WithSellingRegionName("new_zealand"))
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&products, "", "\t")

	err = os.WriteFile("./var/variant_prices.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetCountries(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	countries, err := client.GetCountries( /*printfulsdk.WithLimit(100)*/ /*, printfulsdk.WithTimeout(5*time.Second)*/ )
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&countries, "", "\t")

	err = os.WriteFile("./var/countries.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetTemplates(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	templates, err := client.GetMockupTemplates(403)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&templates, "", "\t")

	err = os.WriteFile("./var/mockup_templates.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetMockupStyles(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	templates, err := client.GetMockupStyles(403)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&templates, "", "\t")

	err = os.WriteFile("./var/mockup_styles.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRequestBody(t *testing.T) {
	opt := printfulsdk.GetOptions(
		printfulsdk.SetURL("https://www.example.com/files/tshirts/example.png"),
	)
	body := printfulsdk.BuildRequestBody(opt, printfulsdk.FileRole, printfulsdk.URL, printfulsdk.Filename, printfulsdk.FileVisible)
	log.Println(body)
}

func TestAddFile(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulsdk.NewPrintfulClient(token)

	file, err := client.AddFile("https://tf2content.loadout.tf/materials/backpack/player/items/sniper/knife_shield.png")
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&file, "", "\t")

	err = os.WriteFile("./var/created_file.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCreateOrder(t *testing.T) {
	recipient := model.Address{
		Name:        "John Smith",
		Address1:    "1 Main St",
		City:        "San Jose",
		CountryCode: "US",
		StateCode:   "CA",
		ZIP:         "95131",
		Email:       "sb-jzssp18153762@personal.example.com",
	}

	items := make([]model.CatalogItem, 0)
	items = append(items, getItem())

	order, err := client.CreateOrder(recipient, items)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&order, "", "\t")

	err = os.WriteFile("./var/created_order.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func getItem() model.CatalogItem {
	item := model.NewCatalogItem()

	item.CatalogVariantID = 19971
	item.Quantity = 1
	item.RetailPrice = "20"
	item.Name = "Test create order"

	placement := model.NewPlacement()
	placement.Placement = "back_large_dtf"
	placement.Technique = "dtfilm"

	layer := model.Layer{}

	layer.Type = "file"
	layer.Url = "https://tf2content.loadout.tf/materials/backpack/weapons/w_models/w_stickybomb_launcher.png"

	placement.Layers = append(placement.Layers, layer)
	item.Placements = append(item.Placements, placement)
	return item
}

func TestGetCategories(t *testing.T) {
	products, err := client.GetCatalogCategories()
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&products, "", "\t")

	err = os.WriteFile("./var/categories.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetSizes(t *testing.T) {
	resp, err := client.Get("https://api.printful.com/v2/catalog-products/785/sizes", nil, context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	response := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&response, "", "\t")

	err = os.WriteFile("./var/sizes.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetPrices(t *testing.T) {
	resp, err := client.Get("https://api.printful.com/v2/catalog-products/785/prices?currency=USD&limit=100", nil, context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	response := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&response, "", "\t")

	err = os.WriteFile("./var/prices_785.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetVariantPrices2(t *testing.T) {
	resp, err := client.Get("https://api.printful.com/v2/catalog-variants/19903/prices?currency=USD&limit=100", nil, context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	response := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
		return
	}
	j, _ := json.MarshalIndent(&response, "", "\t")

	err = os.WriteFile("./var/prices_variant_19903.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetVariants2(t *testing.T) {
	resp, err := client.Get("https://api.printful.com/v2/catalog-products/599/catalog-variants", nil, context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	response := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
		return
	}
	j, _ := json.MarshalIndent(&response, "", "\t")

	err = os.WriteFile("./var/variants_599.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetProduct2(t *testing.T) {
	id := 785
	resp, err := client.Get("https://api.printful.com/v2/catalog-products/"+strconv.Itoa(id), nil, context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	response := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
		return
	}
	j, _ := json.MarshalIndent(&response, "", "\t")

	err = os.WriteFile("./var/product_"+strconv.Itoa(id)+".json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestFetchImage(t *testing.T) {
	img, err := printfulsdk.FetchImage("https://files.cdn.printful.com/m/adidas_space_dyed_polo_shirt/medium/ghost/front/05_adidas_a591_ghost_front_base_whitebg.png")
	if err != nil {
		t.Error(err)
		return
	}

	if img.Bounds().Max.X != 1000 ||
		img.Bounds().Max.Y != 1000 {
		t.Error(errors.New("wrong image size"))
		return
	}

	log.Println(img)
}

func TestGenerateMockup(t *testing.T) {
	inputImage, err := printfulsdk.FetchImage("https://en.wikipedia.org/static/images/icons/wikipedia.png")
	if err != nil {
		t.Error(err)
		return
	}

	mockupTemplates, err := client.GetMockupTemplates(770)
	if err != nil {
		t.Error(err)
		return
	}

	img, err := printfulsdk.GenerateMockup(inputImage, &mockupTemplates[6])
	if err != nil {
		t.Error(err)
		return
	}

	buf := bytes.Buffer{}
	err = png.Encode(&buf, img)
	if err != nil {
		t.Error(err)
		return
	}

	os.WriteFile("test.png", buf.Bytes(), 0666)
}
