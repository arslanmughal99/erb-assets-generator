package main

import (
	"path"
	"flag"
	"imgproc"
	"log"
	"os"
)

// SIZES all linux icons sizes
var SIZES = []int{1024, 512, 256, 128, 96, 64, 48, 32, 24, 16}

func main() {
	var (
		iPath, oPath, name string
	)

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Cannot get current directory: %v", err)
		os.Exit(1)
	}

	flag.StringVar(&name, "n", "icon", "Out icon name")
	flag.StringVar(&oPath, "o", cwd, "Output directory")
	flag.StringVar(&iPath, "i", path.Join(cwd, "icon.png"), "Input icon file")

	flag.Parse()

	img := imgproc.Open(iPath)
	engine := imgproc.New(name, img, SIZES, oPath)
	engine.GenIcons()
}
