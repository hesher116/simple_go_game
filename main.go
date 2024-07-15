package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

// Game репрезентує стан гри.
type Game struct{}

// Update оновлює стан гри. Це викликається кожен кадр (типово 1/60 секунди).
func (g *Game) Update() error {
	// Вихід з гри при натисканні клавіші Q.
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ebiten.ErrRegularTermination
	}
	return nil
}

// Draw рендерить екран гри.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Натисніть Q для виходу")
}

// Layout обробляє розмір вікна гри.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Проста гра на Go")
	if err := ebiten.RunGame(game); err != nil && err != ebiten.ErrRegularTermination {
		log.Fatal(err)
	}
}
