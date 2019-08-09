package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/vector"
)

func drawLines(screen *ebiten.Image) {
	var path vector.Path
	path.Draw(screen, nil)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	drawLines(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 200, 200, 1, "Vector (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
