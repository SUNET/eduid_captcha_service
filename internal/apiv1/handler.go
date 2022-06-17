package apiv1

import (
	"bytes"
	"context"
	"eduid_captcha_service/pkg/model"
	"image/color"

	"github.com/cristalhq/base64"
	"github.com/mcuadros/go-defaults"
	"github.com/steambap/captcha"
)

// RequestGetCaptcha request
type RequestGetCaptcha struct {
	PictureFormat string `json:"picture_format" validate:"omitempty,oneof=jpeg gif" default:"jpeg"`
	PictureSize   struct {
		Width  int `json:"width" default:"400"`
		Height int `json:"height" default:"200"`
	} `json:"picture_size"`
	Noise           float64 `json:"noise" default:"4"`
	BackgroundColor struct {
		R uint8 `json:"R"`
		G uint8 `json:"G"`
		B uint8 `json:"B"`
		A uint8 `json:"A"`
	} `json:"background_color"`
	TextLength  int     `json:"text_length" default:"4"`
	CurveNumber int     `json:"curve_number" default:"2"`
	CharPreset  string  `json:"char_preset" default:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"`
	FontDPI     float64 `json:"font_dpi" default:"72.0"`
}

// ReplyGetCaptcha
type ReplyGetCaptcha struct {
	Text                 string             `json:"text"`
	PictureBase64Encoded string             `json:"picture_base64_encoded"`
	Meta                 *RequestGetCaptcha `json:"meta"`
}

// GetCaptcha handler
func (c *Client) GetCaptcha(ctx context.Context, indata *RequestGetCaptcha) (*ReplyGetCaptcha, error) {
	defaults.SetDefaults(indata)

	opts := func(o *captcha.Options) {
		o.CurveNumber = indata.CurveNumber
		o.TextLength = indata.TextLength
		o.Noise = indata.Noise
		o.CharPreset = indata.CharPreset
		o.BackgroundColor = color.RGBA{
			R: indata.BackgroundColor.R,
			G: indata.BackgroundColor.G,
			B: indata.BackgroundColor.B,
			A: indata.BackgroundColor.A,
		}
		o.FontDPI = indata.FontDPI
	}

	data, err := captcha.New(indata.PictureSize.Width, indata.PictureSize.Height, opts)
	if err != nil {
		return nil, err
	}

	var w bytes.Buffer

	switch indata.PictureFormat {
	case "jpeg":
		if err := data.WriteJPG(&w, nil); err != nil {
			return nil, err
		}
	case "gif":
		if err := data.WriteGIF(&w, nil); err != nil {
			return nil, err
		}
	}

	return &ReplyGetCaptcha{
		Text:                 data.Text,
		PictureBase64Encoded: base64.StdEncoding.EncodeToString(w.Bytes()),
		Meta:                 indata,
	}, nil
}

// Status return status to HAProxy
func (c *Client) Status(ctx context.Context) (*model.Status, error) {
	manyStatus := model.ManyStatus{}

	status := manyStatus.Check()

	return status, nil
}
