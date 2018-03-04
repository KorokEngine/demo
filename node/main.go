package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"korok.io/korok/game"
	"korok.io/korok"
	"korok.io/korok/engi"
	"korok.io/korok/gfx"
	"korok.io/korok/assets"
	"korok.io/korok/hid/input"
)

// A face surround with 4 blocks!
type Block struct {
	engi.Entity
}

func NewBlock() *Block {
	e := korok.Entity.New()
	b := &Block{e}
	korok.Sprite.NewComp(e, nil)
	korok.Transform.NewComp(e)
	return b
}

func (b *Block) SetTexture(tex *gfx.SubTex) {
	korok.Sprite.Comp(b.Entity).SetTexture(tex)
}

func (b *Block) SetPosition(x, y float32) {
	korok.Transform.Comp(b.Entity).SetPosition(mgl32.Vec2{x, y})
}

func (f *Block) SetSize(w, h float32) {
	s := korok.Sprite.Comp(f.Entity)
	s.Width, s.Height = w, h
}


type Face struct {
	engi.Entity
	up, down, left, right *Block
}

func NewFace() *Face {
	e := korok.Entity.New()
	f := &Face{Entity:e}
	korok.Sprite.NewComp(f.Entity, nil)
	korok.Transform.NewComp(f.Entity)
	return f
}

func (f *Face) SetTexture(tex *gfx.SubTex) {
	korok.Sprite.Comp(f.Entity).SetTexture(tex)
}

func (f *Face) SetPosition(x, y float32) {
	korok.Transform.Comp(f.Entity).SetPosition(mgl32.Vec2{x, y})
}


func (f *Face) SetSize(w, h float32) {
	s := korok.Sprite.Comp(f.Entity)
	s.Width, s.Height = w, h
}

func (f *Face) LoadBlock(up, down, left, right *Block) {
	xf := korok.Transform.Comp(f.Entity)
	b1 := korok.Transform.Comp(up.Entity)
	b2 := korok.Transform.Comp(down.Entity)
	b3 := korok.Transform.Comp(left.Entity)
	b4 := korok.Transform.Comp(right.Entity)

	xf.LinkChildren(b1, b2, b3, b4)
	b1.SetPosition(mgl32.Vec2{0, 100})
	b2.SetPosition(mgl32.Vec2{0, -100})
	b3.SetPosition(mgl32.Vec2{-100, 0})
	b4.SetPosition(mgl32.Vec2{100, 0})
}


type MainScene struct {
	face *Face
}

func (m *MainScene) Preload() {
	assets.Texture.Load("assets/face.png")
	assets.Texture.Load("assets/block.png")

	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)
}

func (m *MainScene) Setup(g *game.Game) {
	blockTex := assets.AsSubTexture(assets.Texture.GetTexture("assets/block.png"))
	up, down, left, right := NewBlock(), NewBlock(), NewBlock(), NewBlock()

	up.SetTexture(blockTex); up.SetSize(30, 30)
	down.SetTexture(blockTex); down.SetSize(30, 30)
	left.SetTexture(blockTex); left.SetSize(30, 30)
	right.SetTexture(blockTex); right.SetSize(30, 30)

	faceTex := assets.AsSubTexture(assets.Texture.GetTexture("assets/face.png"))
	face := NewFace()
	face.SetTexture(faceTex)

	face.LoadBlock(up, down, left, right)
	face.SetPosition(240, 160)
	face.SetSize(50 ,50)

	m.face = face
}

func (m *MainScene) Update(dt float32) {
	if dt > 1 {
		return 
	}
	var x, y float32
	if input.Button("up").Down() {
		y = 50
	}
	if input.Button("down").Down() {
		y = -50
	}
	if input.Button("left").Down() {
		x = -50
	}
	if input.Button("right").Down() {
		x = 50
	}

	xf := korok.Transform.Comp(m.face.Entity)
	p := xf.Position()
	x, y = x * dt + p[0] , y * dt + p[1]
	xf.SetPosition(mgl32.Vec2{x, y})
}

func (m *MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	options := &korok.Options{
		Title:"Node System",
		Width:480,
		Height:320,
	}
	korok.Run(options)
}