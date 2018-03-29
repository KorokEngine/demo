package main

import (
	"korok.io/korok/game"
	"korok.io/korok"
	"korok.io/korok/assets"
	"korok.io/korok/effect"
	"korok.io/korok/math/f32"
	"korok.io/korok/math"
)

type MainScene struct {

}

func (*MainScene) Load() {
	assets.Texture.Load("assets/particle.png")
	assets.Texture.Load("assets/block.png")
}

func (*MainScene) OnEnter(g *game.Game) {
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
	gParticle.SetTexture(assets.Texture.Get("assets/particle.png"))
	gXf := korok.Transform.NewComp(gravity)
	gXf.SetPosition(f32.Vec2{240, 160})
}

func (*MainScene) Update(dt float32) {

}

func (*MainScene) OnExit() {
}

func main() {
	options := &korok.Options{
		Title:"ParticleSystem",
		Width:480,
		Height:320,
	}
	korok.Run(options, &MainScene{})
}
