package main

import (
	"fmt"
	"golang-game/renderer"
	"time"
)

func main() {
	renderer.InitalizeLibrary()
	defer renderer.CleanUpLibrary()

	windowName := "main"
	windowHeight := renderer.Type_WindowHeight(800)
	windowWidth := renderer.Type_WindowWidth(800)
	createdRenderer, err := renderer.CreateRenderer(windowName, windowWidth, windowHeight)

	if err != nil {
		fmt.Printf("can't create renderer, %v", err)
	}
	defer createdRenderer.DestroyRenderer()

	rect := renderer.BmpSpriteRect{TopLeftXPos: 0, TopLeftYPos: 0, Width: 100, Height: 100}
	spriteId, err := createdRenderer.CreateBMPFromSpriteMap("assets/test.bmp", rect)
	if err != nil {
		fmt.Printf("failed to create a rect from bmp file, %v", err)
	}

	for !createdRenderer.UserWantsToQuit() {
		err := createdRenderer.Draw()
		sprite := createdRenderer.GetSprite(spriteId)
		if sprite != nil {
			rect.TopLeftXPos += 1
			rect.TopLeftYPos += 1
			rect.TopLeftXPos %= 300
			rect.TopLeftYPos %= 300
			time.Sleep(100 * time.Millisecond)
			sprite.UpdateSpriteRect(rect)
		}
		if err != nil {
			fmt.Printf("failed to render a frame, %v", err)
		}
	}
}
