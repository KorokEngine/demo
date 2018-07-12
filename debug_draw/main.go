package main

import (
	"korok.io/korok/game"
	"korok.io/korok/engi"
	"korok.io/korok/math/f32"
	"korok.io/korok"
	"korok.io/korok/gfx/dbg"
	"korok.io/korok/math"
)

type MainScene struct {
	face engi.Entity
}

func (m *MainScene) OnEnter(g *game.Game) {
}

func (m *MainScene) Update(dt float32) {
	dbg.DrawLine(f32.Vec2{60,60}, f32.Vec2{100, 50})
	dbg.DrawRect(100,50, 20, 20)
	dbg.DrawBorder(130,50, 20, 20, 2)
	dbg.DrawCircle(160, 50, 10)
	dbg.DrawLine(f32.Vec2{50, 200}, f32.Vec2{100, 200})

	center := f32.Vec2{260, 200}
	for i := 0; i < 12; i++ {
		angle := float32(i)/12 * 6.28
		dx := math.Cos(angle)*50 + center[0]
		dy := math.Sin(angle)*50 + center[1]
		dbg.DrawLine(center, f32.Vec2{dx, dy})
	}

	var x, r float32 = 10, 5
	for i := 0; i < 12; i++ {
		dbg.DrawCircle(x, 100, r)
		x += r*2
		r += 5
	}

	dbg.DrawCircle(50, 200, 3)
	dbg.DrawCircle(50, 200, 120)
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