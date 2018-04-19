package main

import (
	"korok.io/korok/game"
	"korok.io/korok/asset"
	"korok.io/korok/gfx"
	"korok.io/korok"
	"korok.io/korok/math/f32"
)


var s_Vertices = []gfx.PosTexColorVertex {
	{ 100.0,  100.0,  1,1, 0xffffffff },
	{-100.0, -100.0,  0,0, 0xffffffff },
	{ 100.0, -100.0,  1,0, 0xffffffff },
	{-100.0,  100.0,  0,1, 0xffffffff },
}

var s_Index = []uint16 {
	3, 1, 2,
	3, 2, 0,
}


var s_mesh *gfx.Mesh
var s_render *gfx.MeshRender

type MainScene struct {

}

func (*MainScene) Load() {
	asset.Texture.Load("face.png")
}

func (*MainScene) OnEnter(g *game.Game) {
	tex2d := asset.Texture.Get("face.png")
	// show mesh comp
	entity := korok.Entity.New()

	comp := korok.Mesh.NewComp(entity)
	mesh := &comp.Mesh

	mesh.SetIndex(s_Index)
	mesh.SetVertex(s_Vertices)
	mesh.Setup()
	mesh.SetTexture(tex2d.Tex())

	xf := korok.Transform.NewComp(entity)
	xf.SetPosition(f32.Vec2{200, 100})
}

func (*MainScene) Update(dt float32) {
}

func (*MainScene) OnExit() {
}

func main() {
	// Run game
	options := &korok.Options{
		Title: "Simple Mesh Rendering",
		Width: 480,
		Height:320,
	}
	korok.Run(options, &MainScene{})
}
