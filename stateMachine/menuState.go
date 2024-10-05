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

// これがメニュー画面のステート
type MenuGameState struct {
	mplusNormalFont font.Face
}

// これがステートが最初に実行されたときに呼び出される関数
func (sm *MenuGameState) Init(
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
func (sm *MenuGameState) Update(stackdeep int) pgfsm.Result {

	/*mキーが入力された場合 メニューを閉じる*/
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		/*ここで現在実行しているメニュー画面のステートマシンを消去します
		「ゲーム画面、メニュー画面」の順でスタックにストックされているので、消去するとスタックの中身は
		「ゲーム画面」となってゲーム画面に戻ります
		*/
		return pgfsm.Result{
			Code:      pgfsm.CodeDelete,
			NextState: nil,
		}
	}
	/*空のpgfsm.Resultを返却することでループを継続します
	pgfsm.Resultを書き換えることで、実行するものを新しいステートに変えたり
	新しいステートをスタックの上に乗せたりすることができます*/
	return pgfsm.Result{}
}

// これはマイフレーム呼び出される描写用の関数です
// このステートが実行されていなくても、スタック上にあれば呼び出されます
func (sm *MenuGameState) Draw(screen *ebiten.Image, stackdeep int) {
	text.Draw(screen, "Menu", sm.mplusNormalFont, 300, 240, color.White)
}
