package mediancut_test

import (
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/alabianca/quantizer/mediancut"
)

func BenchmarkQuantize_8_Colors_MS_3024_4032(b *testing.B) {
	reader, _ := os.Open("../test_images/3024_4032.jpg")
	defer reader.Close()

	img, _ := jpeg.Decode(reader)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		colors := make(color.Palette, 8)
		b.StartTimer()
		_, err := mediancut.Quantize(img, colors, mediancut.MergeSort)
		if err != nil {
			b.Fatalf("expected error to be nil, but got %s", err)
		}
	}
}

func BenchmarkQuantize_8_Colors_MS_2048_1536(b *testing.B) {
	reader, _ := os.Open("../test_images/2048_1536.jpg")
	defer reader.Close()

	img, _ := jpeg.Decode(reader)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		colors := make(color.Palette, 8)
		b.StartTimer()
		_, err := mediancut.Quantize(img, colors, mediancut.MergeSort)
		if err != nil {
			b.Fatalf("expected error to be nil, but got %s", err)
		}
	}
}

func BenchmarkQuantize_8_Colors_QS_3024_4032(b *testing.B) {
	reader, _ := os.Open("../test_images/3024_4032.jpg")
	defer reader.Close()

	img, _ := jpeg.Decode(reader)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		colors := make(color.Palette, 8)
		b.StartTimer()
		_, err := mediancut.Quantize(img, colors, mediancut.QuickSelect)
		if err != nil {
			b.Fatalf("expected error to be nil, but got %s", err)
		}
	}
}

func BenchmarkQuantize_8_Colors_QS_2048_1536(b *testing.B) {
	reader, _ := os.Open("../test_images/2048_1536.jpg")
	defer reader.Close()

	img, _ := jpeg.Decode(reader)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		colors := make(color.Palette, 8)
		b.StartTimer()
		_, err := mediancut.Quantize(img, colors, mediancut.QuickSelect)
		if err != nil {
			b.Fatalf("expected error to be nil, but got %s", err)
		}
	}
}
