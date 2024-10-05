package stateMachine

import (
	"image/color"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// これがタイトル画面のステート
type TitleGameState struct {
	mplusNormalFont font.Face
}

// これがステートが最初に実行されたときに呼び出される関数
func (sm *TitleGameState) Init(
	stackdeep int, /*ここにはこのステートがスタックのどの位置に積まれているかインデックスが入っています*/
) {
	/*ここから Ebitenのフォントの初期化処理*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*ここまで Ebitenのフォントの初期化処理*/
}

// これはマイフレーム呼び出される関数です
// このステートが実行されている時のみ、呼び出されます
func (sm *TitleGameState) Update(
	stackdeep int,
) pgfsm.Result {

	/*sキーが入力された場合*/
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		/*ここでステートマシンを切り替えます
		pgfsm.CodeChangeは現在実行しているステートを
		NextStateに切り替わります
		ここでは現在実行中のタイトル画面のステートからゲーム画面のステートに切り替えています*/
		return pgfsm.Result{
			Code:      pgfsm.CodeChange,
			NextState: &GameMainState{},
		}
	}
	/*空のpgfsm.Resultを返却することでループを継続します
	pgfsm.Resultを書き換えることで、実行するものを新しいステートに変えたり
	新しいステートをスタックの上に乗せたりすることができます*/
	return pgfsm.Result{}
}

// これはマイフレーム呼び出される描写用の関数です
// このステートが実行されていなくても、スタック上にあれば呼び出されます
func (sm *TitleGameState) Draw(screen *ebiten.Image, stackdeep int) {
	text.Draw(screen, "Game Title\nPressing S key,start!", sm.mplusNormalFont, 100, 100, color.White)
}
