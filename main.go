package main

import (
	"fmt"
	"golang-game/game"
	"golang-game/renderer"
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

	gameData := game.GameData{}
	gameData.CreateGameData(&createdRenderer)

	for !createdRenderer.UserWantsToQuit() {
		err := createdRenderer.Draw()
		if err != nil {
			fmt.Printf("failed to render a frame, %v", err)
			continue
		}

		createdRenderer.ClearWindow()
	}
}
