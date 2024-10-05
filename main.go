package main

import (
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func NewGame() *pgfsm.Machine {
	gms := &pgfsm.Machine{}

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	Titlesm := &TitleGameState{}

	/*スタックにタイトル画面のステートを追加します*/
	gms.StateAdd(Titlesm)

	return gms
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(stateMachine.ScreenWidth, stateMachine.ScreenHeight)
	ebiten.SetWindowTitle("Slot Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
