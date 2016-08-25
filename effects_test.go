package bild

import (
	"image"
	"testing"
)

func TestInvert(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xFF,
					0x0, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0xCA, 0x7B, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0x31, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x55, 0x35, 0x84, 0x80, 0x80, 0x80, 0x80, 0xFF,
					0xCE, 0x0, 0xFF, 0xFF, 0xFF, 0x0, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Invert(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("Invert", c.expected, actual))
		}
	}
}

func TestGrayscale(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.Gray
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0xCA, 0x7B, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0x31, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0x0, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xb8, 0x7f,
					0xa8, 0x99,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xCA, 0xAA, 0x7B, 0x80, 0x7f, 0x7f, 0x7A, 0xFF,
					0x71, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xaf, 0x7e,
					0xbb, 0xb3,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Grayscale(c.value)
		if !grayscaleImageEqual(actual, c.expected) {
			t.Error(testFailMessage("Grayscale", c.expected, actual))
		}
	}
}

func TestEdgeDetection(t *testing.T) {
	cases := []struct {
		radius   float64
		value    image.Image
		expected *image.RGBA
	}{
		{
			radius: 0.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x80, 0x80, 0x80, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x80, 0x80, 0x80, 0xff,
				},
			},
		},
		{
			radius: 1.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := EdgeDetection(c.value, c.radius)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("EdgeDetection", c.expected, actual))
		}
	}
}

func TestSobel(t *testing.T) {
	cases := []struct {
		desc     string
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc: "top right corner",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			desc: "bottom right corner",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x40, 0x40, 0x40, 0xff, 0x40, 0x40, 0x40, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Sobel(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("Sobel "+c.desc, c.expected, actual))
		}
	}
}

func TestEmboss(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x2, 0x2, 0x2, 0xff,
					0x80, 0x80, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Emboss(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("Emboss", c.expected, actual))
		}
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		size     int
		value    image.Image
		expected *image.RGBA
	}{
		{
			size: 0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			size: 3,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			size: 2,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF,
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF,
					0x0, 0xFF, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Median(c.value, c.size)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("Sobel", formatImageString(c.expected), formatImageString(actual)))
		}
	}
}

func TestSharpen(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Sharpen(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("Sharpen", c.expected, actual))
		}
	}
}
