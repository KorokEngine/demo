package main

import (
	"korok.io/korok"
	"korok.io/korok/game"
	"korok.io/korok/engi"
	"korok.io/korok/assets"
	"korok.io/korok/gfx"

	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type SpinObject struct {
	engi.Entity
	angle float32
}

func NewSpinObject() *SpinObject {
	e := korok.Entity.New()
	spin := &SpinObject{e,0}
	korok.Sprite.NewComp(e, nil)
	korok.Transform.NewComp(e)
	korok.Script.NewComp(e, spin)
	return spin
}

func (spin *SpinObject) Init() {
}

// 围绕 (240, 160) 的位置旋转, 角速度240，半径 60
func (spin *SpinObject) Update(dt float32) {
	an := dt * 240 / 360 * 6.28
	a := spin.angle + an
	spin.angle = a
	dx := math.Cos(float64(a)) * 60
	dy := math.Sin(float64(a)) * 60
	spin.SetPosition(float32(240 + dx), float32(160 + dy))
}

func (spin *SpinObject) Destroy() {

}

func (spin *SpinObject) SetTexture(tex *gfx.SubTex) {
	if comp := korok.Sprite.Comp(spin.Entity); comp != nil {
		comp.SetTexture(tex)
	}
}

func (spin *SpinObject) SetSize(w, h float32) {
	if comp := korok.Sprite.Comp(spin.Entity); comp != nil {
		comp.Width, comp.Height = w, h
	}
}

func (spin *SpinObject) SetPosition(x, y float32) {
	if comp := korok.Transform.Comp(spin.Entity); comp != nil {
		comp.SetPosition(mgl32.Vec2{x, y})
	}
}

type MainScene struct {
	spin *SpinObject
	angle float32
}

func (*MainScene) Preload() {
	assets.Texture.Load("assets/face.png")
}

func (m *MainScene) Setup(g *game.Game) {
	spin := NewSpinObject()
	spin.SetTexture(assets.AsSubTexture(assets.Texture.GetTexture("assets/face.png")))
	spin.SetSize(30, 30)
	spin.SetPosition(100, 100)
	m.spin = spin
}

func (m *MainScene) Update(dt float32) {

}

func (*MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	options := &korok.Options{
		Title:"Script Demo",
		Width:480,
		Height:320,
	}
	korok.Run(options)
}
