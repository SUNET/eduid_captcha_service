package apiv1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCaptcha(t *testing.T) {
	tts := []struct {
		name string
		have *RequestGetCaptcha
		want *ReplyGetCaptcha
	}{
		{
			name: "default",
			have: &RequestGetCaptcha{},
			want: &ReplyGetCaptcha{
				Meta: &RequestGetCaptcha{
					PictureFormat: "jpeg",
					PictureSize: struct {
						Width  int "json:\"width\" default:\"400\""
						Height int "json:\"height\" default:\"200\""
					}{
						Width:  400,
						Height: 200,
					},
					Noise: 4,
					BackgroundColor: struct {
						R uint8 "json:\"R\""
						G uint8 "json:\"G\""
						B uint8 "json:\"B\""
						A uint8 "json:\"A\""
					}{
						R: 0,
						G: 0,
						B: 0,
						A: 0,
					},
					TextLength:  4,
					CurveNumber: 2,
					CharPreset:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
					FontDPI:     72.0,
				},
			},
		},
		{
			name: "gif",
			have: &RequestGetCaptcha{
				PictureFormat: "gif",
			},
			want: &ReplyGetCaptcha{
				Meta: &RequestGetCaptcha{
					PictureFormat: "gif",
					PictureSize: struct {
						Width  int "json:\"width\" default:\"400\""
						Height int "json:\"height\" default:\"200\""
					}{
						Width:  400,
						Height: 200,
					},
					Noise: 4,
					BackgroundColor: struct {
						R uint8 "json:\"R\""
						G uint8 "json:\"G\""
						B uint8 "json:\"B\""
						A uint8 "json:\"A\""
					}{
						R: 0,
						G: 0,
						B: 0,
						A: 0,
					},
					TextLength:  4,
					CurveNumber: 2,
					CharPreset:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
					FontDPI:     72.0,
				},
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			client := mockClient(t)

			got, err := client.GetCaptcha(context.TODO(), tt.have)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Meta, got.Meta)
		})
	}
}
