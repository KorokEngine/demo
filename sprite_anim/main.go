package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/asset"
	"korok.io/korok/engi"
	"korok.io/korok/hid/input"
	"korok.io/korok/gfx"
	"korok.io/korok/math/f32"
)

type MainScene struct {
	hero engi.Entity
	g *game.Game
}

func (*MainScene) Load() {
	asset.Texture.LoadAtlasIndexed("hero.png", 52, 72, 4, 3)
}

func (m *MainScene) OnEnter(g *game.Game) {
	// get animation system...

	// input control
	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)

	hero := korok.Entity.New()

	// SpriteComp
	korok.Sprite.NewComp(hero).SetSize(50 ,50)
	korok.Transform.NewComp(hero).SetPosition(f32.Vec2{240, 160})

	fb := korok.Flipbook.NewComp(hero)
	fb.SetRate(.2)

	m.hero = hero
	{
		at, _ := asset.Texture.Atlas("hero.png")
		frames := [12]gfx.Tex2D{}
		for i := 0; i < 12; i++ {
			frames[i], _ = at.GetByIndex(i)
		}
		g.SpriteEngine.NewAnimation("hero.down", frames[0:3], true)
		g.SpriteEngine.NewAnimation("hero.left", frames[3:6], true)
		g.SpriteEngine.NewAnimation("hero.right", frames[6:9], true)
		g.SpriteEngine.NewAnimation("hero.top", frames[9:12], true)
	}

	// default
	fb.Play("hero.down")
}

func (m *MainScene) Update(dt float32) {
	speed := f32.Vec2{0, 0}

	// 根据上下左右，执行不同的帧动画
	if input.Button("up").JustPressed() {
		korok.Flipbook.Comp(m.hero).Play("hero.top")
	}
	if input.Button("down").JustPressed() {
		korok.Flipbook.Comp(m.hero).Play("hero.down")
	}
	if input.Button("left").JustPressed() {
		korok.Flipbook.Comp(m.hero).Play("hero.left")
	}
	if input.Button("right").JustPressed() {
		korok.Flipbook.Comp(m.hero).Play("hero.right")
	}

	scalar := float32(3)
	if input.Button("up").Down() {
		speed[1] = scalar
	}
	if input.Button("down").Down() {
		speed[1] = -scalar
	}
	if input.Button("left").Down() {
		speed[0] = -scalar
	}
	if input.Button("right").Down() {
		speed[0] = scalar
	}

	xf := korok.Transform.Comp(m.hero)

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