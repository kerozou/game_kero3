package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kerozou/kero3/stateMachine/stateMachine"
)

const ()

func main() {
	Rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	game := stateMachine.NewGame(Rand)

	ebiten.SetWindowSize(stateMachine.ScreenWidth, stateMachine.ScreenHeight)
	ebiten.SetWindowTitle("Slot Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
