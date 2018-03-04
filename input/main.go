package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/assets"
	"korok.io/korok/engi"
	"korok.io/korok/hid/input"

	"github.com/go-gl/mathgl/mgl32"
)

type MainScene struct {
	face engi.Entity
}

func (*MainScene) Preload() {
	assets.Texture.Load("assets/face.png")

	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)
}

// 输入
func (m *MainScene) Setup(g *game.Game) {
	id, tex := assets.Texture.GetTexture("assets/face.png")

	face := korok.Entity.New()
	comp := korok.Sprite.NewComp(face, assets.AsSubTexture(id, tex))
	comp.Width = 50
	comp.Height = 50

	faceXF := korok.Transform.NewComp(face)
	faceXF.SetPosition(mgl32.Vec2{10, 100})

	m.face = face
}

func (m *MainScene) Update(dt float32) {
	speed := mgl32.Vec2{0, 0}
	if input.Button("up").Down() {
		speed[1] = 5
	}
	if input.Button("down").Down() {
		speed[1] = -5
	}
	if input.Button("left").Down() {
		speed[0] = -5
	}
	if input.Button("right").Down() {
		speed[0] = 5
	}

	xf := korok.Transform.Comp(m.face)
	p := xf.Position()
	xf.SetPosition(mgl32.Vec2{p[0]+speed[0], p[1]+speed[1]})
}

func (*MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	options := &korok.Options{
		Title:"Input Controller",
		Width:480,
		Height:320,
	}
	korok.Run(options)
}
