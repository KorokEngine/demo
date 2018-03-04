package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/assets"
	"github.com/go-gl/mathgl/mgl32"

	"math/rand"
)

type MainScene struct {
}

func (*MainScene) Preload() {
	assets.Texture.Load("assets/face.png")
	assets.Texture.Load("assets/block.png")
}

func (m *MainScene) Setup(g *game.Game) {
	// show blocks
	id, tex := assets.Texture.GetTexture("assets/block.png")
	for i := 0; i < 800; i++ {
		entity := korok.Entity.New()
		c := korok.Sprite.NewComp(entity, assets.AsSubTexture(id, tex))
		c.Width = 20
		c.Height = 20

		xf := korok.Transform.NewComp(entity)

		x := float32(rand.Intn(480))
		y := float32(rand.Intn(200)) + 120
		xf.SetPosition(mgl32.Vec2{x, y})
	}

	// show face
	{
		id, tex := assets.Texture.GetTexture("assets/face.png")
		face := korok.Entity.New()
		sprite := korok.Sprite.NewComp(face, assets.AsSubTexture(id, tex))
		sprite.SetSize(50, 50)

		xf := korok.Transform.NewComp(face)
		xf.SetPosition(mgl32.Vec2{100, 20})
	}
}

func (m *MainScene) Update(dt float32) {
}

func (*MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	// Run game
	options := &korok.Options{
		Title: "Sprite Rendering",
		Width: 480,
		Height:320,
	}
	korok.Run(options)
}