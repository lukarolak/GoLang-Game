package renderer

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	window renderWindow
}

type Type_WindowWidth uint32
type Type_WindowHeight uint32

func CreateRenderer(windowName string, windowWidth Type_WindowWidth, windowHeight Type_WindowHeight) (error, Renderer) {
	renderer := Renderer{}
	var err error = nil
	err, renderer.window = createWindow(windowName, windowWidth, windowHeight)

	if err != nil {
		return fmt.Errorf("can't create window, %w", err), renderer
	}

	return nil, renderer
}

func (renderer Renderer) createWindow(windowWidth Type_WindowWidth, windowHeight Type_WindowHeight) (error, *sdl.Window) {
	if math.MaxInt32 < windowWidth {
		return fmt.Errorf("specified window width (%u) not supported", windowWidth), nil
	}

	if math.MaxInt32 < windowHeight {
		return fmt.Errorf("specified window height (%u) not supported", windowHeight), nil
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("can't create window, %w", err), nil
	}

	return nil, window
}

func (renderer Renderer) DestroyRenderer() {
	renderer.window.Destroy()
}

func (renderer Renderer) Draw() error {
	err := renderer.window.Draw()

	if err != nil {
		return fmt.Errorf("can't draw to the window, %w", err)
	}

	return nil
}

func (renderer Renderer) UserWantsToQuit() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
			break
		}
	}
	return false
}

func (renderer *Renderer) CreateRectFromBMP(pathToBMP string) (error, Type_RectId) {
	return renderer.window.CreateRectFromBMP(pathToBMP)
}
