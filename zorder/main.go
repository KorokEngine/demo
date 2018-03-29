package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/asset"
	"korok.io/korok/engi"
	"korok.io/korok/hid/input"
	"korok.io/korok/math/f32"
)

type MainScene struct {
	face engi.Entity
}

func (*MainScene) Load() {
	asset.Texture.Load("assets/face.png")
	asset.Texture.Load("assets/block.png")
}

func (m *MainScene) OnEnter(g *game.Game) {
	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)

	input.RegisterButton("Order", input.Q)

	tex := asset.Texture.Get("assets/block.png")

	// blocks
	for i := 0; i < 8; i++ {
		block := korok.Entity.New()
		sprite := korok.Sprite.NewCompX(block, tex)
		sprite.SetSize(30, 30)
		sprite.SetZOrder(int16(i))

		xf := korok.Transform.NewComp(block)
		x := float32(i * 40)
		y := float32(200)
		xf.SetPosition(f32.Vec2{x, y})
	}

	// face
	{
		face := korok.Entity.New()

		tex := asset.Texture.Get("assets/face.png")
		sprite := korok.Sprite.NewCompX(face, tex)
		sprite.SetSize(50, 50)

		blockXF := korok.Transform.NewComp(face)
		blockXF.SetPosition(f32.Vec2{100, 20})

		m.face = face
	}
}

var index = 0

var orderList = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func (m *MainScene) Update(dt float32) {
	speed := mgl32.Vec2{0, 0}
	if input.Button("up").Down() {
		speed[1] = 10
	}
	if input.Button("down").Down() {
		speed[1] = -10
	}
	if input.Button("left").Down() {
		speed[0] = -10
	}
	if input.Button("right").Down() {
		speed[0] = 10
	}

	if input.Button("Order").JustPressed() {
		korok.Sprite.Comp(m.face).SetZOrder(orderList[index%10]); index ++
	}

	xf := korok.Transform.Comp(m.face)

	x := xf.Position()[0] + speed[0]
	y := xf.Position()[1] + speed[1]

	xf.SetPosition(f32.Vec2{x, y})
}

func (*MainScene) OnExit() {
}

func main() {
	// Run game
	options := &korok.Options{
		Title: "Hello, Korok Engine",
		Width: 480,
		Height:320,
	}
	korok.Run(options, &MainScene{})
}