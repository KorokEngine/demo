package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/asset"
	"korok.io/korok/engi"
	"korok.io/korok/hid/input"
	"korok.io/korok/math/f32"
	"korok.io/korok/effect"
	"korok.io/korok/math"
	"korok.io/korok/gfx/font"
	"korok.io/korok/gfx"
)

type MainScene struct {
	face engi.Entity
}

func (*MainScene) Load() {
	asset.Texture.Load("face.png")
	asset.Texture.Load("block.png")
	asset.Texture.Load("particle.png")
	// font
	asset.Font.LoadTrueType("font1", "OCRAEXT.TTF", font.ASCII(64))
}

func (m *MainScene) OnEnter(g *game.Game) {
	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)

	input.RegisterButton("Order", input.Q)

	tex := asset.Texture.Get("block.png")
	fnt, _ := asset.Font.Get("font1")

	// face variable z-order 0-9
	{
		face := korok.Entity.New()

		tex := asset.Texture.Get("face.png")
		sprite := korok.Sprite.NewCompX(face, tex)
		sprite.SetSize(50, 50)

		blockXF := korok.Transform.NewComp(face)
		blockXF.SetPosition(f32.Vec2{200, 80})

		m.face = face
	}

	// blocks z-order: [0, 7]
	for i := 0; i < 8; i++ {
		block := korok.Entity.New()
		sprite := korok.Sprite.NewCompX(block, tex)
		sprite.SetSize(30, 30)
		sprite.SetZOrder(int16(i))

		xf := korok.Transform.NewComp(block)
		x := float32(i * 40) + 80
		y := float32(200)
		xf.SetPosition(f32.Vec2{x, y})
	}

	// text z-order: 6
	{
		hello := korok.Entity.New()
		text := korok.Text.NewComp(hello)
		text.SetFont(fnt)
		text.SetFontSize(18)
		text.SetColor(gfx.Red)
		text.SetText("Hello World")
		text.SetZOrder(6)

		xf := korok.Transform.NewComp(hello)
		xf.SetPosition(f32.Vec2{240, 240})
		xf.RotateBy(.57)
	}

	// particle z-order:0
	{
		cfg := &effect.GravityConfig{
			Config:effect.Config {
				Max:1024,
				Rate:10,
				Duration:math.MaxFloat32,
				Life:effect.Var{40.1, 0.4},
				Size:effect.Range{effect.Var{10 ,5}, effect.Var{20, 5}},
				X:effect.Var{0, 0}, Y:effect.Var{0, 0},
				A: effect.Range{effect.Var{1, 0}, effect.Var{0, 0}},
			},
			Speed: effect.Var{70, 10},
			Angel: effect.Var{math.Radian(90), math.Radian(30)},
			Gravity:f32.Vec2{0, -10},
		}
		gravity := korok.Entity.New()
		gParticle := korok.ParticleSystem.NewComp(gravity)
		gParticle.SetSimulator(effect.NewGravitySimulator(cfg))
		gParticle.SetTexture(asset.Texture.Get("particle.png"))
		xf := korok.Transform.NewComp(gravity)
		xf.SetPosition(f32.Vec2{40, 160})
	}

}

var index = 0

var orderList = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func (m *MainScene) Update(dt float32) {
	speed := f32.Vec2{0, 0}
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