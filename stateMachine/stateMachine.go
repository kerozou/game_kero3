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

// これがゲーム画面のステート
type GameMainState struct {
	mplusNormalFont font.Face
}

// これがステートが最初に実行されたときに呼び出される関数
func (sm *GameMainState) Init(
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
// つまりメニューを開いている間は、ゲーム画面のUpdate関数が実行されません
func (sm *GameMainState) Update(stackdeep int) pgfsm.Result {
	/*mキーが入力された場合 メニューを開く*/
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		/*ここで現在実行しているゲーム画面の上にメニュー画面のステートをのせます
		「ゲーム画面」の順でスタックにストックされているので、追加するとスタックの中身は
		「ゲーム画面、メニュー画面」となってメニュー画面の処理に移ります
		*/
		return pgfsm.Result{
			Code:      pgfsm.CodeAdd,
			NextState: &MenuGameState{},
		}
	}

	/*空のpgfsm.Resultを返却することでループを継続します
	pgfsm.Resultを書き換えることで、実行するものを新しいステートに変えたり
	新しいステートをスタックの上に乗せたりすることができます*/
	return pgfsm.Result{}
}

// これはマイフレーム呼び出される描写用の関数です
// このステートが実行されていなくても、スタック上にあれば呼び出されます
// つまりメニューを開いている間も、ゲーム画面のdraw関数が実行されます
func (sm *GameMainState) Draw(screen *ebiten.Image, stackdeep int) {
	text.Draw(screen, "Game Main", sm.mplusNormalFont, 200, 100, color.White)
}

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
