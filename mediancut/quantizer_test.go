package mediancut_test

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"os"
	"testing"

	"github.com/alabianca/quantizer/mediancut"
)

func TestQuantizer(t *testing.T) {
	reader, _ := os.Open("../test_images/3024_4032.jpg")
	defer reader.Close()

	img, _ := jpeg.Decode(reader)

	colors := make([]mediancut.Point, 8)
	_, err := mediancut.Quantize(img, colors, mediancut.QuickSelect)
	if err != nil {
		t.Fatalf("expected error to be nil, but got %s", err)
	}

	expected := []mediancut.Point{
		{
			Red:   8,
			Green: 8,
			Blue:  25,
		},
		{
			Red:   35,
			Green: 38,
			Blue:  33,
		},
		{
			Red:   52,
			Green: 54,
			Blue:  60,
		},
		{
			Red:   70,
			Green: 85,
			Blue:  122,
		},
		{
			Red:   108,
			Green: 79,
			Blue:  90,
		},
		{
			Red:   144,
			Green: 124,
			Blue:  102,
		},
		{
			Red:   110,
			Green: 128,
			Blue:  182,
		},
		{
			Red:   151,
			Green: 170,
			Blue:  209,
		},
	}

	for i, c := range colors {
		verifyColor(c, expected[i], t)
	}

}

func verifyColor(c mediancut.Point, expected mediancut.Point, t *testing.T) {
	if c.Red != expected.Red {
		t.Fatalf("expected Red %d, but got %d", c.Red, expected.Red)
	}
	if c.Green != expected.Green {
		t.Fatalf("expected Green %d, but got %d", c.Green, expected.Green)
	}
	if c.Blue != expected.Blue {
		t.Fatalf("expected Blue %d, but got %d", c.Blue, expected.Blue)
	}
}

func imageToBase64String(img image.Image, t *testing.T) string {
	buf := bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buf, img, nil); err != nil {
		t.Fatalf("jpeg encode error %s, but expected no error", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
