package kero3

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	reelCount    = 3
	symbolCount  = 10 // 10行に変更
	symbolWidth  = 64
	symbolHeight = 64
)

var reelSymbols = [10][3]int{
	{0, 1, 0},
	{1, 2, 0},
	{1, 1, 1},
	{2, 3, 1},
	{1, 4, 2},
	{3, 5, 1},
	{4, 0, 3},
	{5, 2, 4},
	{0, 1, 5},
	{2, 4, 0},
}

type Game struct {
	rand           *rand.Rand
	reels          [reelCount][3]int // 各リールがSymbolCountのindexを持つ(0~9) 配列のi
	isFlashing     bool
	flashStartTime time.Time
	spinning       bool
	finished       bool
	spinCount      int
	spinTarget     [reelCount]int
	spinSpeed      [reelCount]float64
	audioContext   *audio.Context
	se1            []byte // うぉ
	se2            []byte // おはようございます
	reelImage2     *ebiten.Image
	barImage       *ebiten.Image
}

func NewGame(Rand *rand.Rand) *Game {
	// Create the game
	game := &Game{rand: Rand}
	game.Init()
	return game
}

func (g *Game) Init() {
	// 1.リール画像の読み込み
	img, err := Embed.Open("reel2.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(img)
	if err != nil {
		log.Fatal(err)
	}
	g.reelImage2 = ebiten.NewImageFromImage(p)

	// 2.真ん中のバーの表示
	img2, err := Embed.Open("bar.png")
	if err != nil {
		log.Fatal(err)
	}
	p2, err := png.Decode(img2)
	if err != nil {
		log.Fatal(err)
	}
	g.barImage = ebiten.NewImageFromImage(p2)

	// 音設定
	g.audioContext = audio.NewContext(44100)
	g.se1, err = openSE("1.mp3")
	if err != nil {
		log.Fatalf("failed to load sound effect 1: %v", err)
	}
	g.se2, err = openSE("2.mp3")
	if err != nil {
		log.Fatalf("failed to load sound effect 1: %v", err)
	}

	// リールの初期化
	for i := 0; i < reelCount; i++ {
		for j := 0; j < 3; j++ {
			g.reels[i][j] = i
		}
	}

}

func openSE(name string) ([]byte, error) {
	file, err := Embed.Open(name)
	if err != nil {
		panic(err)
	}
	src, err2 := mp3.DecodeWithoutResampling(file)
	if err2 != nil {
		panic(err2)
	}
	return io.ReadAll(src)
}

func (g *Game) Update() error {
	// スペースキーがちょうど押されたかどうかをチェック
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !g.spinning {
		g.spinning = true
		g.finished = false
		g.isFlashing = false
		g.flashStartTime = time.Now()
		g.spinCount = 0
		g.spinTarget[0] = g.rand.Intn(16) + 8  // 8~12
		g.spinTarget[1] = g.rand.Intn(45) + 13 // 13~18
		g.spinTarget[2] = g.rand.Intn(37) + 15 // 15~21
		g.spinSpeed[0] = 5.0
		g.spinSpeed[1] = 4.0
		g.spinSpeed[2] = 3.0

		// Goルーチンでスピン処理を実行
		go g.spinReels()
	}

	// フラッシュの継続時間をチェック
	if g.isFlashing && time.Since(g.flashStartTime) > 100*time.Millisecond {
		g.isFlashing = false
	}

	return nil
}

func (g *Game) spinReels() {
	done := make(chan bool)

	// Goルーチンでスピン処理を実行
	go func() {
		// 上段を確定する
		for j := 0; j < reelCount; j++ {
			g.reels[0][j] = (g.spinTarget[j] + g.reels[0][j]) % symbolCount
		}

		// 中段, 下段も追従させる
		for j := 0; j < reelCount; j++ {
			g.reels[1][j] = (g.reels[0][j] + 1) % symbolCount
			g.reels[2][j] = (g.reels[1][j] + 1) % symbolCount
		}
		g.spinning = false

		// スピンが終了したことを通知
		done <- true
	}()

	// スピン処理が完了するのを待つ
	<-done

	// スピンが終了したらfinished処理を実行
	g.finished = true
	g.checkReels()
}

func (g *Game) checkReels() {
	done := make(chan bool)
	ok := true
	g.isFlashing = true

	// Goルーチンでforループを実行
	go func() {
		for i := 0; i < reelCount; i++ {
			fmt.Printf("[1][%d] = %d, [1][0] = %d\n", i, reelSymbols[g.reels[1][i]][i], reelSymbols[g.reels[1][0]][0])
			if reelSymbols[g.reels[1][i]][i] != reelSymbols[g.reels[1][0]][0] {
				ok = false
				g.isFlashing = false
			}
		}
		done <- true
	}()

	// forループが終了するのを待つ
	<-done

	// 後続の処理
	if ok {
		sePlayer := g.audioContext.NewPlayerFromBytes(g.se2)
		sePlayer.SetVolume(0.3) // 音量を30%に設定
		sePlayer.Play()
	} else {
		sePlayer := g.audioContext.NewPlayerFromBytes(g.se1)
		sePlayer.SetVolume(0.1) // 音量を10%に設定
		sePlayer.Play()
	}
	g.finished = false
	fmt.Printf("end\n")
}

func (g *Game) Draw(screen *ebiten.Image) {
	// フラッシュ
	if g.isFlashing {
		// 画面を白色で塗りつぶす
		screen.Fill(color.RGBA{255, 255, 255, 255})
	} else {
		// 通常の描画
		screen.Fill(color.RGBA{0, 0, 0, 255})
	}

	centerX := (ScreenWidth - reelCount*symbolWidth - 230) / 2
	centerY := (ScreenHeight - 3*symbolHeight) / 2

	// バー表示
	opbar := &ebiten.DrawImageOptions{}
	opbar.GeoM.Translate(float64(centerX-50), float64(centerY+87))
	opbar.ColorScale.Scale(1, 1, 1, 0.2) // 透明度50%
	screen.DrawImage(g.barImage, opbar)

	// 回転停止後描画
	if !g.spinning {
		// リールの回転が終了したときにスプライトの位置をピクセル単位で調整
		for i := 0; i < reelCount; i++ {
			for j := 0; j < 3; j++ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(centerX+j*symbolWidth), float64(centerY+i*symbolHeight))
				symbolIndex := g.reels[i][j]
				screen.DrawImage(g.reelImage2.SubImage(image.Rect(symbolWidth*j, symbolIndex*symbolHeight, symbolWidth*(j+1), (symbolIndex+1)*symbolHeight)).(*ebiten.Image), op)
			}
		}
		ebitenutil.DebugPrint(screen, "Press SPACE to spin")

		// 2行目の各マスのスプライトの値を画面の右下に描画
		for i := 0; i < reelCount; i++ {
			for j := 0; j < 3; j++ {
				text := fmt.Sprintf("%d", reelSymbols[g.reels[i][j]][j])
				ebitenutil.DebugPrintAt(screen, text, ScreenWidth-240+symbolWidth*j, ScreenHeight-320+(symbolHeight+10)*i)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
