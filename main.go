package main

import (
	"fmt"
	"golang-game/renderer"
)

func main() {
	renderer.InitalizeLibrary()
	defer renderer.CleanUpLibrary()

	windowName := "main"
	windowHeight := renderer.Type_WindowHeight(800)
	windowWidth := renderer.Type_WindowWidth(800)
	err, renderer := renderer.CreateRenderer(windowName, windowWidth, windowHeight)

	if err != nil {
		fmt.Printf("can't create renderer, %w", err)
	}
	defer renderer.DestroyRenderer()

	err, _ = renderer.CreateRectFromBMP("assets/test.bmp")
	if err != nil {
		fmt.Printf("failed to create a rect from bmp file, %w", err)
	}

	for renderer.UserWantsToQuit() == false {
		err := renderer.Draw()
		if err != nil {
			fmt.Printf("failed to render a frame, %w", err)
		}
	}
}
