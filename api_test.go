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

func TestGetProducts(t *testing.T) {
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
		//time.Sleep(10 * time.Millisecond)

		go func() {
			defer wg.Done()
			client.GetCountries()
			done = done + 1
			//log.Println(done)
		}()
	}

	wg.Wait()

	err = client.GetCatalogProducts()
	if err != nil {
		t.Error(err)
		return
	}
}
