package renderer

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type renderSurface struct {
	Surface  *sdl.Surface
	BMPRects []*bmpRect
}

func (window renderWindow) createSurface() (error, renderSurface) {
	surface := renderSurface{}

	var err error = nil
	surface.Surface, err = window.Window.GetSurface()
	if err != nil {
		return fmt.Errorf("can't create surface, %w", err), surface
	}

	return nil, surface
}

type Type_RectId uint32

func (surface *renderSurface) CreateRectFromBMP(pathToBMP string) (error, Type_RectId) {
	err, rect := createBMPRect(pathToBMP)

	if err != nil {
		return fmt.Errorf("can't create rect from BMP, %w", err), Type_RectId(0)
	}

	for i := 0; i < len(surface.BMPRects); i++ {
		if surface.BMPRects[i] == nil {
			surface.BMPRects[i] = &rect
			return nil, Type_RectId(i)
		}
	}

	surface.BMPRects = append(surface.BMPRects, &rect)
	return nil, Type_RectId(len(surface.BMPRects) - 1)
}

func (surface renderSurface) Draw() error {
	for index, rect := range surface.BMPRects {
		if rect != nil {
			err := surface.drawBMPRect(*rect)
			if err != nil {
				return fmt.Errorf("can't draw bmp with id %u", Type_RectId(index))
			}
		}
	}

	return nil
}

func (surface renderSurface) drawBMPRect(rect bmpRect) error {
	err := rect.Surface.Blit(nil, surface.Surface, &sdl.Rect{X: 0, Y: 0, W: rect.Surface.W, H: rect.Surface.H})
	if err != nil {
		return fmt.Errorf("can't draw BMP rect, %w", err)
	}

	return err
}

func (surface renderSurface) Destroy() {
	for _, rect := range surface.BMPRects {
		if rect != nil {
			rect.DestroyBMPRect()
		}
	}
}
