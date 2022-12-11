package renderer

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bmpRect struct {
	Surface *sdl.Surface
}

type BmpSpriteRect struct {
	TopLeftXPos uint32
	TopLeftYPos uint32
	Width       uint32
	Height      uint32
}

func (rect BmpSpriteRect) validateForSDLUse() error {
	if math.MaxInt32 < rect.TopLeftXPos {
		return fmt.Errorf("topLeftXPos out of bounds")
	}
	if math.MaxInt32 < rect.TopLeftYPos {
		return fmt.Errorf("topLeftYPos out of bounds")
	}
	if math.MaxInt32 < rect.Width {
		return fmt.Errorf("width out of bounds")
	}
	if math.MaxInt32 < rect.Height {
		return fmt.Errorf("height out of bounds")
	}

	return nil
}

func createBMPFromFile(pathToBMP string) (bmpRect, error) {
	bmp := bmpRect{}

	var err error = nil
	bmp.Surface, err = sdl.LoadBMP(pathToBMP)
	if err != nil {
		return bmp, fmt.Errorf("can't load BMP from path %s, %w", pathToBMP, err)
	}
	return bmp, nil
}

func (bmp bmpRect) DestroyBMPRect() {
	bmp.Surface.Free()
}

func createBMPRectFromSpriteMap(pathToBMP string, bmpSpriteRect BmpSpriteRect) (bmpRect, error) {
	bmp, err := createBMPFromFile(pathToBMP)
	if err != nil {
		return bmp, fmt.Errorf("can't create BMP rect, %w", err)
	}
	err = bmpSpriteRect.validateForSDLUse()
	if err != nil {
		return bmp, fmt.Errorf("can't use provided bmp sprite rect, %w", err)
	}

	bmp.Surface.SetClipRect(&sdl.Rect{
		X: int32(bmpSpriteRect.TopLeftXPos),
		Y: int32(bmpSpriteRect.TopLeftYPos),
		W: int32(bmpSpriteRect.Width),
		H: int32(bmpSpriteRect.Height)})

	return bmp, nil
}

func (rect *bmpRect) UpdateSpriteRect(newSpriteRect BmpSpriteRect) error {
	err := newSpriteRect.validateForSDLUse()
	if err != nil {
		return fmt.Errorf("can't use provided bmp sprite rect, %w", err)
	}

	rect.Surface.SetClipRect(&sdl.Rect{
		X: int32(newSpriteRect.TopLeftXPos),
		Y: int32(newSpriteRect.TopLeftYPos),
		W: int32(newSpriteRect.Width),
		H: int32(newSpriteRect.Height)})

	return nil
}
