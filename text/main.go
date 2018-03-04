package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/assets"

	"github.com/go-gl/mathgl/mgl32"
)

type MainScene struct {
}

func (*MainScene) Preload() {
	assets.Font.LoadBitmap("dft",
		"assets/font/font.png",
		"assets/font/font.json")
}

func (*MainScene) Setup(g *game.Game) {
	font := assets.Font.GetFont("dft")

	// show "Hello world"
	entity := korok.Entity.New()

	text := korok.Text.NewComp(entity)
	text.SetFont(font)
	text.SetText("Hello Korok!")

	xf := korok.Transform.NewComp(entity)
	xf.SetPosition(mgl32.Vec2{30, 100})
}

func (*MainScene) Update(dt float32) {

}

func (*MainScene) Name() string {
	return "Main"
}

func main()  {
	korok.PushScene(&MainScene{})

	options := korok.Options{
		Title:"Text Rendering",
		Width:480,
		Height:320,
	}
	korok.Run(&options)
}
