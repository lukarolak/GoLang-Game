package renderer

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type renderSurface struct {
	Surface       *sdl.Surface
	bmpCollection renderBMPCollection
}

func (window renderWindow) createSurface() (renderSurface, error) {
	surface := renderSurface{}

	var err error = nil
	surface.Surface, err = window.Window.GetSurface()
	if err != nil {
		return surface, fmt.Errorf("can't create surface, %w", err)
	}

	return surface, nil
}

func (surface *renderSurface) CreateBMPFromFile(pathToBMP string) (Type_SpriteId, error) {
	bmp, err := createBMPFromFile(pathToBMP)

	if err != nil {
		return Type_SpriteId(0), fmt.Errorf("can't create rect from BMP, %w", err)
	}

	return surface.bmpCollection.AddBMP(bmp), nil
}

func (surface *renderSurface) CreateBMPFromSpriteMap(pathToBMP string, bmpSpriteRect BmpSpriteRect) (Type_SpriteId, error) {
	rect, err := createBMPRectFromSpriteMap(pathToBMP, bmpSpriteRect)

	if err != nil {
		return Type_SpriteId(0), fmt.Errorf("can't create rect from BMP, %w", err)
	}

	return surface.bmpCollection.AddBMP(rect), nil
}

func (surface renderSurface) Draw() error {
	for index, rect := range surface.bmpCollection.BMPRects {
		if rect != nil {
			err := surface.drawBMPRect(*rect)
			if err != nil {
				return fmt.Errorf("can't draw bmp with id %d", Type_SpriteId(index))
			}
		}
	}

	return nil
}

func (surface renderSurface) drawBMPRect(rect bmpRect) error {
	err := rect.Surface.Blit(&rect.Surface.ClipRect, surface.Surface, &sdl.Rect{X: 0, Y: 0, W: rect.Surface.W, H: rect.Surface.H})
	if err != nil {
		return fmt.Errorf("can't draw BMP rect, %w", err)
	}

	return err
}

func (surface renderSurface) Destroy() {
	surface.bmpCollection.Destroy()
}

func (surface *renderSurface) GetSprite(spriteId Type_SpriteId) *bmpRect {
	return surface.bmpCollection.GetSprite(spriteId)
}
