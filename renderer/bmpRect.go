package renderer

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type bmpRect struct {
	Surface *sdl.Surface
}

func createBMPRect(pathToBMP string) (error, bmpRect) {
	bmp := bmpRect{}

	var err error = nil
	bmp.Surface, err = sdl.LoadBMP(pathToBMP)
	if err != nil {
		return fmt.Errorf("can't load BMP from path %s, %w", pathToBMP, err), bmp
	}

	return nil, bmp
}

func (bmp bmpRect) DestroyBMPRect() {
	bmp.Surface.Free()
}
