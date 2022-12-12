package renderer

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type renderWindow struct {
	Window  *sdl.Window
	Surface renderSurface
}

func createWindow(windowName string, windowWidth Type_WindowWidth, windowHeight Type_WindowHeight) (renderWindow, error) {
	window := renderWindow{}

	if math.MaxInt32 < windowWidth {
		return window, fmt.Errorf("specified window width (%d) not supported", windowWidth)
	}

	if math.MaxInt32 < windowHeight {
		return window, fmt.Errorf("specified window height (%d) not supported", windowHeight)
	}

	var err error = nil
	window.Window, err = sdl.CreateWindow(windowName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)

	if err != nil {
		return window, fmt.Errorf("can't create window, %w", err)
	}

	window.Surface, err = window.createSurface()
	if err != nil {
		return window, fmt.Errorf("can't create surface, %w", err)
	}

	return window, nil
}

func (window renderWindow) Destroy() {
	window.Surface.Destroy()
	window.Window.Destroy()
}

func (window renderWindow) Draw() error {
	err := window.Surface.Draw()
	if err != nil {
		return fmt.Errorf("can't draw on the surface, %w", err)
	}
	window.Window.UpdateSurface()

	return nil
}

func (window renderWindow) ClearWindow() error {
	err := window.Surface.Clear()
	if err != nil {
		return fmt.Errorf("can't clear window surface, %w", err)
	}

	return nil
}

func (window *renderWindow) CreateBMPFromFile(pathToBMP string) (Type_SpriteId, error) {
	return window.Surface.CreateBMPFromFile(pathToBMP)
}

func (window *renderWindow) CreateBMPFromSpriteMap(pathToBMP string, bmpSpriteRect BmpSpriteRect) (Type_SpriteId, error) {
	return window.Surface.CreateBMPFromSpriteMap(pathToBMP, bmpSpriteRect)
}

func (window *renderWindow) GetSprite(spriteId Type_SpriteId) *bmpRect {
	return window.Surface.GetSprite(spriteId)
}
