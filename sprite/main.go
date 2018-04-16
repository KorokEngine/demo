package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/asset"
	"korok.io/korok/math/f32"

	"math/rand"
)

type MainScene struct {
}

func (*MainScene) Load() {
	asset.Texture.Load("face.png")
	asset.Texture.Load("block.png")
}

func (m *MainScene) OnEnter(g *game.Game) {
	// show blocks
	tex := asset.Texture.Get("block.png")
	for i := 0; i < 800; i++ {
		entity := korok.Entity.New()
		korok.Sprite.NewCompX(entity,tex).SetSize(20, 20)

		xf := korok.Transform.NewComp(entity)

		x := float32(rand.Intn(480))
		y := float32(rand.Intn(200)) + 120
		xf.SetPosition(f32.Vec2{x, y})
	}

	// show face
	{
		tex := asset.Texture.Get("face.png")
		face := korok.Entity.New()
		korok.Sprite.NewCompX(face, tex).SetSize(50 ,50)

		xf := korok.Transform.NewComp(face)
		xf.SetPosition(f32.Vec2{100, 20})
	}
}

func (m *MainScene) Update(dt float32) {
}

func (*MainScene) OnExit() {
}

func main() {
	// Run game
	options := &korok.Options{
		Title: "Sprite Rendering",
		Width: 480,
		Height:320,
	}
	korok.Run(options, &MainScene{})
}