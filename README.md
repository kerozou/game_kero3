# game_kero3
スロットマシーンのゲーム

## 概要
golang, ebitengine の練習用に作ってるものになります。


## ステートマシンについて
/main.go を実行すると、/stateMachine/stateMachine.go のNewGame()が呼び出されます。
以降、下記のステートを遷移します
- TitleGameState: タイトル画面
- MenuGameState: メニュー画面
- GameMainState: ゲームのメイン画面

上記のようなステートを追加して遷移図を作ることで、シーン遷移が実現できるのと、スタック化もできるのでシーンを積み上げることができるはずです