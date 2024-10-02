package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kerozou/stateMachine/stateMachine"
)

const ()

func main() {
	game := stateMachine.NewGame()

	ebiten.SetWindowSize(stateMachine.ScreenWidth, stateMachine.ScreenHeight)
	ebiten.SetWindowTitle("Slot Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
