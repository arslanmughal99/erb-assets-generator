package encoders

import (
	"github.com/Kodeworks/golang-image-ico"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/toukii/icns"
	"image"
	"io"
)

// ICNSEncoder returns Icns Encoder
func ICNSEncoder() imgio.Encoder {
	return func(w io.Writer, img image.Image) error {
		return icns.Encode(w, img)
	}
}

// IcoEncoder returns Ico Encoder
func IcoEncoder() imgio.Encoder {
	return func(w io.Writer, img image.Image) error {
		return ico.Encode(w, img)
	}
}
