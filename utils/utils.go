package utils

import (
	"os"
	"image"
	_ "image/jpeg"
	"image/draw"
	"github.com/wieku/glhf"
	"log"
)

func LoadImage(path string) (*image.NRGBA, error) {
	file, err := os.Open(path)
	log.Println("Loading texture: ", file.Name())
	if err != nil {
		log.Println("er1")
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("er2")
		return nil, err
	}
	bounds := img.Bounds()
	nrgba := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)
	return nrgba, nil
}

func LoadTexture(path string) (*glhf.Texture, error) {
	img, err := LoadImage(path)
	if err == nil {
		tex := glhf.NewTexture(
			img.Bounds().Dx(),
			img.Bounds().Dy(),
			4,
			true,
			img.Pix,
		)

		tex.Begin()
		tex.SetWrap(glhf.CLAMP_TO_EDGE)
		tex.End()

		return tex, nil
	}
	return nil, err
}
