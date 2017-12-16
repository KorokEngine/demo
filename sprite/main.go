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

	face := korok.Entity.New()
	comp := korok.Sprite.NewComp(face, assets.AsSubTexture(id, tex))
	comp.Width = 50
	comp.Height = 50

	faceXF := korok.Transform.NewComp(face)
	faceXF.Position = mgl32.Vec2{10, 100}

	for i := 0; i < 800; i++ {
		face := korok.Entity.New()
		c := korok.Sprite.NewComp(face, assets.AsSubTexture(id, tex))
		c.Width = 20
		c.Height = 20

		faceXF := korok.Transform.NewComp(face)

		x := float32(rand.Intn(480))
		y := float32(rand.Intn(200)) + 120
		faceXF.Position = mgl32.Vec2{x, y}
	}

	// show face
	bid, btex := assets.Texture.GetTexture("assets/face.png")
	block := korok.Entity.New()
	comp = korok.Sprite.NewComp(block, assets.AsSubTexture(bid, btex))
	comp.Width = 50
	comp.Height = 50

	blockXF := korok.Transform.NewComp(block)
	blockXF.Position = mgl32.Vec2{100, 20}
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