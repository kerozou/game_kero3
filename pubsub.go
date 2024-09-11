package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type EventType string

const (
	EventTypeKeyPress EventType = "KeyPress"
)

type Event struct {
	Type EventType
	Data interface{}
}

type Subscriber func(event Event)

type EventBus struct {
	subscribers map[EventType][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]Subscriber),
	}
}

func (bus *EventBus) Subscribe(eventType EventType, subscriber Subscriber) {
	bus.subscribers[eventType] = append(bus.subscribers[eventType], subscriber)
}

func (bus *EventBus) Publish(event Event) {
	if subs, found := bus.subscribers[event.Type]; found {
		for _, sub := range subs {
			sub(event)
		}
	}
}

type Game struct {
	eventBus *EventBus
}

func NewGame() *Game {
	eventBus := NewEventBus()
	game := &Game{eventBus: eventBus}

	// キー押下イベントの購読者を追加
	eventBus.Subscribe(EventTypeKeyPress, func(event Event) {
		key := event.Data.(ebiten.Key)
		fmt.Printf("Key pressed: %v\n", key)
	})

	return game
}

func (g *Game) Update() error {
	// スペースキーが押されたらイベントを発行
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.eventBus.Publish(Event{Type: EventTypeKeyPress, Data: ebiten.KeySpace})
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // 黒色で画面を塗りつぶす
	ebitenutil.DebugPrint(screen, "Press SPACE to trigger event")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pub/Sub Example")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
