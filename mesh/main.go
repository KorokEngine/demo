package main

import (
	"korok.io/korok/game"
	"korok.io/korok/assets"
	"korok.io/korok/gfx"
	"korok.io/korok"

	"github.com/go-gl/mathgl/mgl32"
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

func (*MainScene) Preload() {
	assets.Texture.Load("assets/face.png")
}

func (*MainScene) Setup(g *game.Game) {
	// show mesh comp
	entity := korok.Entity.New()

	comp := korok.Mesh.NewComp(entity)
	mesh := &comp.Mesh

	mesh.SetIndex(s_Index)
	mesh.SetVertex(s_Vertices)
	mesh.Setup()
	mesh.TextureId, _ = assets.Texture.GetTexture("assets/face.png")

	xf := korok.Transform.NewComp(entity)
	xf.Position = mgl32.Vec2{200, 100}
}

func (*MainScene) Update(dt float32) {
}

func (*MainScene) Name() string {
	return "main"
}

func main() {
	korok.PushScene(&MainScene{})
	// Run game
	options := &korok.Options{
		Title: "Simple Mesh Rendering",
		Width: 480,
		Height:320,
	}
	korok.Run(options)
}
