package mediancut_test

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/alabianca/quantizer/mediancut"
)

func TestQuantizer(t *testing.T) {
	reader, _ := os.Open("../test_images/3024_4032.jpg")
	defer reader.Close()

	img, _ := jpeg.Decode(reader)

	colors := make(color.Palette, 8)
	_, err := mediancut.Quantize(img, colors, mediancut.QuickSelect)
	if err != nil {
		t.Fatalf("expected error to be nil, but got %s", err)
	}

	expected := []color.NRGBA64{
		{R: 2170, G: 2245, B: 6551, A: 65535},
		{R: 9167, G: 9776, B: 8680, A: 65535},
		{R: 13362, G: 13844, B: 15580, A: 65535},
		{R: 17930, G: 21998, B: 31464, A: 65535},
		{R: 27682, G: 20339, B: 23124, A: 65535},
		{R: 37046, G: 31858, B: 26350, A: 65535},
		{R: 28252, G: 32785, B: 46721, A: 65535},
		{R: 38702, G: 43612, B: 53714, A: 65535},
	}

	for i, c := range colors {
		verifyColor(c, expected[i], t)
	}

}

func verifyColor(c color.Color, expected color.Color, t *testing.T) {
	r, g, b, _ := c.RGBA()
	re, ge, be, _ := expected.RGBA()
	if r != re {
		t.Fatalf("expected Red %d, but got %d", r, re)
	}
	if g != ge {
		t.Fatalf("expected Green %d, but got %d", g, ge)
	}
	if b != be {
		t.Fatalf("expected Blue %d, but got %d", b, be)
	}
}

func imageToBase64String(img image.Image, t *testing.T) string {
	buf := bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buf, img, nil); err != nil {
		t.Fatalf("jpeg encode error %s, but expected no error", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
