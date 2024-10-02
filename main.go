package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kerozou/kero3/stateMachine/stateMachine"
)

const ()

func main() {
	gms := stateMachine.pgfsm.Machine{}

	Titlesm := &TitleGameState{}

	/*スタックにタイトル画面のステートを追加します*/
	gms.StateAdd(Titlesm)

	ebiten.SetWindowSize(stateMachine.ScreenWidth, stateMachine.ScreenHeight)
	ebiten.SetWindowTitle("Slot Game")
	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}
