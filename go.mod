module github.com/kerozou/kero3

go 1.22.6

require (
	github.com/PenguinCabinet/pgfsm v0.0.7
	github.com/hajimehoshi/ebiten/v2 v2.8.0
	golang.org/x/image v0.20.0
)

require (
	github.com/ebitengine/gomobile v0.0.0-20240911145611-4856209ac325 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/oto/v3 v3.3.0 // indirect
	github.com/ebitengine/purego v0.8.0 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	github.com/kerozou/stateMachine/stateMachine v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)

replace (
	github.com/kerozou/kero3/kero3 => ./kero3/
	github.com/kerozou/stateMachine/titleState => ./titleState
	github.com/kerozou/stateMachine/menuState => ./menuState
	github.com/kerozou/stateMachine/gameMainState => ./gameMainState
)
