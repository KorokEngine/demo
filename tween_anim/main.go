package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/assets"
	"korok.io/korok/engi"
	"korok.io/korok/anim/tween"
	"korok.io/korok/anim/tween/ease"
	"korok.io/korok/math/f32"
)

type MainScene struct {
	hero engi.Entity
	g *game.Game
	en *tween.Engine
}

func (*MainScene) Load() {

	assets.Texture.Load("assets/face.png")
}

func (m *MainScene) OnEnter(g *game.Game) {
	m.en = g.TweenEngine

	// texture
	tex := assets.Texture.Get("assets/face.png")

	// ease functions
	funcs := []ease.Function {
		ease.Linear,
		ease.OutCirc,
		ease.OutBounce,
		ease.OutElastic,
		ease.OutBack,
		ease.OutCubic,
	}

	for i := range funcs {
		entity := korok.Entity.New()
		korok.Sprite.NewCompX(entity, tex).SetSize(30, 30)
		xf := korok.Transform.NewComp(entity)
		ii := i

		animator := m.en.NewAnimator()
		animator.SetValue(10, 240).SetDuration(2).SetFunction(funcs[i]).SetRepeat(tween.RepeatInfinite).OnUpdate(func (f, v float32) {
			xf.SetPosition(f32.Vec2{v, 50 + 30 *float32(ii)})
		}).Start()
	}
}

func (m *MainScene) Update(dt float32) {
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