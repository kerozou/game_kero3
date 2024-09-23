package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kerozou/kero3/kero3"
)

const (
)
func main() {
	Rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	game := kero3.NewGame(Rand)

	ebiten.SetWindowSize(kero3.ScreenWidth, kero3.ScreenHeight)
	ebiten.SetWindowTitle("Slot Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
