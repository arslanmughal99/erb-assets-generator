package imgproc

import (
	"encoders"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"image"
	"log"
	"os"
	"path"
	"strconv"
)

// Icons hold all the methods and config to process
type Icons struct {
	name    string      // name of output icon
	out     string      // out output folder default to `pwd`
	sizes   []int       // sizes
	rawIcon image.Image // rawIcon is the unprocessed icon image
}

// New returns new `Icons`
func New(name string, img image.Image, sizes []int, out string) Icons {
	return Icons{name, out, sizes, img}
}

// GenIcons generte all the icons
func (i *Icons) GenIcons() {
	i.genIco()
	i.genIcns()
	i.genPng()
}

// `genIcns` generate icns file
func (i *Icons) genIcns() {
	outPath := path.Join(i.out, i.name+".icns")
	save(outPath, i.rawIcon, encoders.ICNSEncoder())
}

// `genIco` generate ico file
func (i *Icons) genIco() {
	outPath := path.Join(i.out, i.name+".ico")
	save(outPath, i.rawIcon, encoders.IcoEncoder())
}

// `genPng` generate all linux png icons files
func (i *Icons) genPng() {
	if err := os.MkdirAll(path.Join(path.Dir(i.out), "/icons"), os.ModeDir.Perm()); err != nil {
		log.Fatalf("Fail to create linux icons dir: %v", err)
		os.Exit(1)
	}

	for _, size := range i.sizes {
		rImg := resize(i.rawIcon, size, size)
		outPath := path.Join(i.out, "/icons/", strconv.Itoa(size)+"x"+strconv.Itoa(size)+".png")
		save(outPath, rImg, imgio.PNGEncoder())
	}
}

// Open image from given path and returns `image.Image`
// `path` icon path
func Open(p string) image.Image {
	img, err := imgio.Open(p)
	if err != nil {
		log.Fatalf("Error While opening image: %v", err)
		os.Exit(1)
	}
	return img
}

// `resize` the given image according to given `width`, `height` and `Encoder`
func resize(img image.Image, height, width int) *image.RGBA {
	rezImg := transform.Resize(img, width, height, transform.Linear)
	return rezImg
}

// `save` icon
func save(path string, img image.Image, enc imgio.Encoder) {
	if err := imgio.Save(path, img, enc); err != nil {
		log.Fatalf("Fail to save icon: %v", err)
		os.Exit(1)
	}
}
