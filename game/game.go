package game

import (
	"fmt"
	"golang-game/renderer"
)

type GameData struct {
}

func (data GameData) CreateGameData(render *renderer.Renderer) {
	rect := renderer.BmpSpriteRect{TopLeftXPos: 97, TopLeftYPos: 1, Width: 14, Height: 14}

	for x := uint32(0); x < 3; x++ {
		for y := uint32(0); y < 3; y++ {
			spriteId, err := render.CreateBMPFromSpriteMap("assets/Tile.bmp", rect)
			if err != nil {
				fmt.Printf("failed to create a rect from bmp file, %v", err)
			}
			render.GetSprite(spriteId).UpdateSpritePosition(14*x, 14*y)
		}
	}
}
