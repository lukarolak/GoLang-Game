package renderer

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	window renderWindow
}

type Type_WindowWidth uint32
type Type_WindowHeight uint32

func CreateRenderer(windowName string, windowWidth Type_WindowWidth, windowHeight Type_WindowHeight) (Renderer, error) {
	renderer := Renderer{}
	var err error = nil
	renderer.window, err = createWindow(windowName, windowWidth, windowHeight)

	if err != nil {
		return renderer, fmt.Errorf("can't create window, %w", err)
	}

	return renderer, nil
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
		}
	}
	return false
}

func (renderer *Renderer) CreateBMPFromFile(pathToBMP string) (Type_SpriteId, error) {
	return renderer.window.CreateBMPFromFile(pathToBMP)
}

func (renderer *Renderer) CreateBMPFromSpriteMap(pathToBMP string, bmpSpriteRect BmpSpriteRect) (Type_SpriteId, error) {
	return renderer.window.Surface.CreateBMPFromSpriteMap(pathToBMP, bmpSpriteRect)
}

func (renderer Renderer) GetSprite(spriteId Type_SpriteId) *bmpRect {
	return renderer.window.GetSprite(spriteId)
}
