package printfulsdk

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/icza/gox/imagex/colorx"
	"golang.org/x/image/draw"
)

const (
	TemplatePositioningOverlay    string = "overlay"
	TemplatePositioningBackground string = "background"
)

func FetchImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	config, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if config.Width > 20000 || config.Height > 20000 {
		return nil, fmt.Errorf("image is too large: %dx%d", config.Width, config.Height)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func GenerateMockup(i image.Image, t *model.MockupTemplates) (image.Image, error) {
	if i == nil {
		return nil, errors.New("image is nil")
	}
	if t == nil {
		return nil, errors.New("template is nil")
	}

	mockup := image.NewNRGBA(image.Rect(0, 0, int(t.TemplateWidth), int(t.TemplateHeight)))

	c := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	if t.BackgroundColor != "" {
		var err error
		c, err = colorx.ParseHexColor(t.BackgroundColor)
		if err != nil {
			return nil, fmt.Errorf("failed to parse color: %s", t.BackgroundColor)
		}
	}

	u := image.NewUniform(c)
	draw.Draw(mockup, mockup.Bounds(), u, image.Pt(0, 0), draw.Over)

	if t.BackgroundURL != "" {
		img, err := FetchImage(t.BackgroundURL)
		if err != nil {
			return nil, err
		}
		//draw.Draw(mockup, mockup.Bounds(), img, image.Pt(0, 0), draw.Src)
		draw.BiLinear.Scale(mockup, mockup.Bounds(), img, img.Bounds(), draw.Over, nil)
	}

	if t.TemplatePositioning == TemplatePositioningBackground {
		if t.ImageURL != "" {
			img, err := FetchImage(t.ImageURL)
			if err != nil {
				return nil, err
			}
			//draw.Draw(mockup, mockup.Bounds(), img, image.Pt(0, 0), draw.Src)
			draw.BiLinear.Scale(mockup, mockup.Bounds(), img, img.Bounds(), draw.Over, nil)
		}
	}

	draw.BiLinear.Scale(mockup, image.Rect(int(t.PrintAreaLeft), int(t.PrintAreaTop), int(t.PrintAreaLeft+t.PrintAreaWidth), int(t.PrintAreaTop+t.PrintAreaHeight)), i, i.Bounds(), draw.Over, nil)

	if t.TemplatePositioning == TemplatePositioningOverlay {
		img, err := FetchImage(t.ImageURL)
		if err != nil {
			return nil, err
		}

		draw.BiLinear.Scale(mockup, mockup.Bounds(), img, img.Bounds(), draw.Over, nil)
	}

	return mockup, nil
}
