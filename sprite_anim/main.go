package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/assets"
	"korok.io/korok/engi"
	"korok.io/korok/hid/input"
	"korok.io/korok/gfx"
	"korok.io/korok/anim/frame"
)

type MainScene struct {
	hero engi.Entity
	g *game.Game
	as *frame.Engine
}

func (*MainScene) Preload() {
	assets.Texture.Load("assets/face.png")
	assets.Texture.Load("assets/block.png")
	assets.Texture.Load("assets/hero.png")
}

func (m *MainScene) Setup(g *game.Game) {
	// get animation system...
	m.as = g.SpriteEngine

	// input control
	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)

	hero := korok.Entity.New()

	// SpriteComp
	id, tex := assets.Texture.GetTexture("assets/hero.png")
	sprite := korok.Sprite.NewComp(hero, nil)
	sprite.Width = 50
	sprite.Height = 50


	xf := korok.Transform.NewComp(hero)
	xf.SetPosition(mgl32.Vec2{240, 160})

	m.hero = hero
	{
		frames := [3]gfx.SubTex{}
		w, h := tex.Width, tex.Height
		fill := func(frames *[3]gfx.SubTex, x, y float32) {
			for i := 0; i < 3; i++ {
				x1 := (x+52*float32(i))/w
				y1 := y/h
				x2 := (x+52*float32(i+1))/w
				y2 := (y+72)/h
				frames[i].TexId = id
				frames[i].Width = 52
				frames[i].Height = 72
				frames[i].Region = gfx.Region{x1, y1, x2, y2}
			}
		}

		// 哪些动画是需要更新的呢？哪些又不需要更新呢？
		// 如何管理？
		fill(&frames, 0, 0)
		m.as.NewAnimation("hero.down", frames[:], true)
		fill(&frames, 0, 72)
		m.as.NewAnimation("hero.left", frames[:], true)
		fill(&frames, 0, 72*2)
		m.as.NewAnimation("hero.right", frames[:], true)
		fill(&frames, 0, 72*3)
		m.as.NewAnimation("hero.top", frames[:], true)
	}

	// default
	m.as.Of(m.hero).Rate(.2).Play("hero.down")
}

func (m *MainScene) Update(dt float32) {
	speed := mgl32.Vec2{0, 0}

	// 根据上下左右，执行不同的帧动画
	if input.Button("up").JustPressed() {
		m.as.Of(m.hero).Rate(.2).Play("hero.top")
	}
	if input.Button("down").JustPressed() {
		m.as.Of(m.hero).Rate(.2).Play("hero.down")
	}
	if input.Button("left").JustPressed() {
		m.as.Of(m.hero).Rate(.2).Play("hero.left")
	}
	if input.Button("right").JustPressed() {
		m.as.Of(m.hero).Rate(.2).Play("hero.right")
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
	xf.SetPosition(mgl32.Vec2{x, y})
}

func (*MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	// Run game
	options := &korok.Options{
		Title: "Hello, Korok Engine",
		Width: 480,
		Height:320,
	}
	korok.Run(options)
}