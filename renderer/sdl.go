package renderer

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func InitalizeLibrary() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("can't initalize sdl, %w", err)
	}
	return nil
}

// Call to clean up resources opened by the SDL library
func CleanUpLibrary() {
	sdl.Quit()
}
