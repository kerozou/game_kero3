package main

import (
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kerozou/kero3/kero3"
	"github.com/kerozou/stateMachine/stateMachine"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func toTitleState() *pgfsm.Machine {
	gms := &pgfsm.Machine{}

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	Titlesm := &stateMachine.TitleState{}

	/*スタックにタイトル画面のステートを追加します*/
	gms.StateAdd(Titlesm)

	return gms
}

func main() {
	game := toTitleState()

	ebiten.SetWindowSize(kero3.ScreenWidth, kero3.ScreenHeight)
	ebiten.SetWindowTitle("Slot Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
