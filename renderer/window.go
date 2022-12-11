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

func createWindow(windowName string, windowWidth Type_WindowWidth, windowHeight Type_WindowHeight) (error, renderWindow) {
	window := renderWindow{}

	if math.MaxInt32 < windowWidth {
		return fmt.Errorf("specified window width (%u) not supported", windowWidth), window
	}

	if math.MaxInt32 < windowHeight {
		return fmt.Errorf("specified window height (%u) not supported", windowHeight), window
	}

	var err error = nil
	window.Window, err = sdl.CreateWindow(windowName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("can't create window, %w", err), window
	}

	err, window.Surface = window.createSurface()
	if err != nil {
		return fmt.Errorf("can't create surface, %w", err), window
	}

	return nil, window
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

func (window *renderWindow) CreateRectFromBMP(pathToBMP string) (error, Type_RectId) {
	return window.Surface.CreateRectFromBMP(pathToBMP)
}
