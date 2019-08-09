package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func drawEbitenText(screen *ebiten.Image) {
	var path vector.Path

	// E
	path.MoveTo(60, 20)
	path.LineTo(20, 20)
	path.LineTo(20, 60)
	path.LineTo(60, 60)
	path.MoveTo(20, 40)
	path.LineTo(60, 40)

	// B
	path.MoveTo(110, 20)
	path.LineTo(80, 20)
	path.LineTo(80, 60)
	path.LineTo(110, 60)
	path.LineTo(120, 50)
	path.LineTo(110, 40)
	path.LineTo(120, 30)
	path.LineTo(110, 20)
	path.LineTo(80, 20)
	path.MoveTo(80, 40)
	path.LineTo(110, 40)

	// I
	path.MoveTo(140, 20)
	path.LineTo(140, 60)

	// T
	path.MoveTo(160, 20)
	path.LineTo(200, 20)
	path.MoveTo(180, 20)
	path.LineTo(180, 60)

	// E
	path.MoveTo(260, 20)
	path.LineTo(220, 20)
	path.LineTo(220, 60)
	path.LineTo(260, 60)
	path.MoveTo(220, 40)
	path.LineTo(260, 40)

	// N
	path.MoveTo(280, 60)
	path.LineTo(280, 20)
	path.LineTo(320, 60)
	path.LineTo(320, 20)

	op := &vector.DrawPathOptions{}
	op.LineWidth = 8
	op.StrokeColor = color.RGBA{0xdb, 0x56, 0x20, 0xff}
	path.Draw(screen, op)
}

func drawLines(screen *ebiten.Image) {
	var path vector.Path
	path.MoveTo(20, 80+float32(counter%320)/4)
	path.LineTo(60, 120)
	path.LineTo(20, 160-float32(counter%320)/4)

	op := &vector.DrawPathOptions{}
	op.LineWidth = 16
	op.StrokeColor = color.White
	path.Draw(screen, op)
}

var counter = 0

func update(screen *ebiten.Image) error {
	counter++
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	drawEbitenText(screen)
	drawLines(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Vector (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
