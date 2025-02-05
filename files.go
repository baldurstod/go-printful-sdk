package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) AddFile(url string, opts ...requestOption) (*model.File, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.url = url

	body := BuildRequestBody(opt, FileRole, URL, Filename, FileVisible)

	u := "https://api.printful.com/v2/files"
	resp, err := c.Post(u, nil, body, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("post returned an error in AddFile: %w", err)
	}

	response := &responses.AddFileResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response")
	}

	return &response.Data, nil
}
