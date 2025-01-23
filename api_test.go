package printfulsdk_test

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	printfulsdk "github.com/baldurstod/go-printful-sdk"
)

type Config struct {
	Printful `json:"printful"`
}

type Printful struct {
	AccessToken string `json:"access_token"`
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

	countries, err := client.GetTemplates(403)
	if err != nil {
		t.Error(err)
		return
	}

	j, _ := json.MarshalIndent(&countries, "", "\t")

	err = os.WriteFile("./var/templates.json", j, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}
