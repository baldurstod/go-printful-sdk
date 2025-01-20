package printfulapi_test

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	printfulapi "github.com/baldurstod/go-printful-api"
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

	client := printfulapi.NewPrintfulClient(token)

	client.GetCountries()
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup

	var done = 0
	for i := 1; i <= 130; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			client.GetCountries( /*printfulapi.WithTimeout(1 * time.Second)*/ )
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

	client := printfulapi.NewPrintfulClient(token)

	products, err := client.GetCatalogProducts(printfulapi.WithLimit(100) /*, printfulapi.WithTimeout(5*time.Second)*/)
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

func TestGetCountries(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, err := getAuthToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := printfulapi.NewPrintfulClient(token)

	countries, err := client.GetCountries(printfulapi.WithLimit(100) /*, printfulapi.WithTimeout(5*time.Second)*/)
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
