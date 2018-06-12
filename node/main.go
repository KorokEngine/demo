package main

import (
	"korok.io/korok/game"
	"korok.io/korok"
	"korok.io/korok/engi"
	"korok.io/korok/gfx"
	"korok.io/korok/asset"
	"korok.io/korok/hid/input"
	"korok.io/korok/math/f32"
)

// A face surround with 4 blocks!
type Block struct {
	engi.Entity
}

func NewBlock() Block {
	e := korok.Entity.New()
	b := Block{e}
	korok.Sprite.NewComp(e)
	korok.Transform.NewComp(e)
	return b
}

func (b *Block) SetTexture(tex gfx.Tex2D) {
	korok.Sprite.Comp(b.Entity).SetSprite(tex)
}

func (b *Block) SetPosition(x, y float32) {
	korok.Transform.Comp(b.Entity).SetPosition(f32.Vec2{x, y})
}

func (b *Block) SetSize(w, h float32) {
	korok.Sprite.Comp(b.Entity).SetSize(w, h)
}


type Face struct {
	engi.Entity
	up, down, left, right Block
}

func NewFace() *Face {
	e := korok.Entity.New()
	f := &Face{Entity:e}
	korok.Sprite.NewComp(f.Entity)
	korok.Transform.NewComp(f.Entity)
	return f
}

func (f *Face) SetTexture(tex gfx.Tex2D) {
	korok.Sprite.Comp(f.Entity).SetSprite(tex)
}

func (f *Face) SetPosition(x, y float32) {
	korok.Transform.Comp(f.Entity).SetPosition(f32.Vec2{x, y})
}

func (f *Face) SetSize(w, h float32) {
	korok.Sprite.Comp(f.Entity).SetSize(w, h)
}

func (f *Face) LoadBlock(up, down, left, right Block) {
	xf := korok.Transform.Comp(f.Entity)
	b1 := korok.Transform.Comp(up.Entity)
	b2 := korok.Transform.Comp(down.Entity)
	b3 := korok.Transform.Comp(left.Entity)
	b4 := korok.Transform.Comp(right.Entity)

	xf.LinkChildren(b1, b2, b3, b4)
	b1.SetPosition(f32.Vec2{0, 100})
	b2.SetPosition(f32.Vec2{0, -100})
	b3.SetPosition(f32.Vec2{-100, 0})
	b4.SetPosition(f32.Vec2{100, 0})
}


type MainScene struct {
	face *Face
}

func (m *MainScene) Load() {
	asset.Texture.Load("face.png")
	asset.Texture.Load("block.png")

	input.RegisterButton("up", input.ArrowUp)
	input.RegisterButton("down", input.ArrowDown)
	input.RegisterButton("left", input.ArrowLeft)
	input.RegisterButton("right", input.ArrowRight)
}

func (m *MainScene) OnEnter(g *game.Game) {
	blockTex := asset.Texture.Get("block.png")
	up, down, left, right := NewBlock(), NewBlock(), NewBlock(), NewBlock()

	up.SetTexture(blockTex); up.SetSize(30, 30)
	down.SetTexture(blockTex); down.SetSize(30, 30)
	left.SetTexture(blockTex); left.SetSize(30, 30)
	right.SetTexture(blockTex); right.SetSize(30, 30)

	faceTex := asset.Texture.Get("face.png")
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
	xf.SetPosition(f32.Vec2{x, y})
}

func (m *MainScene) OnExit() {
}

func main() {
	options := &korok.Options{
		Title:"Node System",
		Width:480,
		Height:320,
	}
	korok.Run(options, &MainScene{})
}