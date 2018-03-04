package main

import (
	"korok.io/korok/game"
	"korok.io/korok"
	"korok.io/korok/assets"
	"korok.io/korok/effect"

	"github.com/go-gl/mathgl/mgl32"
)

type MainScene struct {

}

func (*MainScene) Preload() {
	assets.Texture.Load("assets/particle.png")
	assets.Texture.Load("assets/block.png")
}

func (*MainScene) Setup(g *game.Game) {
	cfg := &effect.GravityConfig{
		Config:effect.Config {
			Max:10240,
			Duration:0,
			Life:effect.Var{40.1, 0.4},
			Size:effect.Range{effect.Var{10 ,5}, effect.Var{20, 5}},
			X:effect.Var{0, 0}, Y:effect.Var{0, 0},
			A: effect.Range{effect.Var{1, 0}, effect.Var{0, 0}},
		},
		Velocity: [2]effect.Var{{-30, 80}, {10, 30}},
		Gravity:mgl32.Vec2{0, -10},
	}
	gravity := korok.Entity.New()
	gParticle := korok.ParticleSystem.NewComp(gravity)
	gParticle.SetSimulator(effect.NewGravitySimulator(cfg))
	gParticle.SetTexture(assets.AsSubTexture(assets.Texture.GetTexture("assets/particle.png")))
	gXf := korok.Transform.NewComp(gravity)
	gXf.SetPosition(mgl32.Vec2{240, 160})
}

func (*MainScene) Update(dt float32) {

}

func (*MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	options := &korok.Options{
		Title:"ParticleSystem",
		Width:480,
		Height:320,
	}
	korok.Run(options)
}
