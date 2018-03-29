package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/asset"
	"korok.io/korok/math/f32"
)

type MainScene struct {
}

func (*MainScene) Load() {
	asset.Font.LoadBitmap("dft",
		"assets/font/font.png",
		"assets/font/font.json")
}

func (*MainScene) OnEnter(g *game.Game) {
	font := asset.Font.GetFont("dft")

	// show "Hello world"
	entity := korok.Entity.New()

	text := korok.Text.NewComp(entity)
	text.SetFont(font)
	text.SetText("Hello Korok!")

	xf := korok.Transform.NewComp(entity)
	xf.SetPosition(f32.Vec2{30, 100})
}

func (*MainScene) Update(dt float32) {

}

func (*MainScene) OnExit() {
}

func main()  {
	options := korok.Options{
		Title:"Text Rendering",
		Width:480,
		Height:320,
	}
	korok.Run(&options, &MainScene{})
}
