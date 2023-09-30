package img

import (
	"image"
	"io"

	"golang.org/x/image/draw"
)

// OpenImage reads an image from an io.Reader and returns an image object
func OpenImage(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}

// ScaleImage rescales an image given an image rectagle and returns rescaled image object
func ScaleImage(src image.Image, rect image.Rectangle) image.Image {
	dst := image.NewRGBA(rect)
	scale := draw.CatmullRom
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}
