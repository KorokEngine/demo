package main

import (
	"korok.io/korok/gui"
	"korok.io/korok/hid/input"
	"korok.io/korok/assets"
	"korok.io/korok/game"
	"korok.io/korok/gfx"
	"korok.io/korok"

	"log"
	"fmt"
)

type MainScene struct {
	face gfx.Tex2D
	slide float32

	normal, pressed gfx.Tex2D
	showbutton bool
}

func (m *MainScene) Preload() {
	assets.Texture.Load("assets/face.png")
	assets.Texture.Load("assets/block.png")
	assets.Texture.Load("assets/particle.png")
	assets.Font.LoadBitmap("asc", "assets/font/font.png", "assets/font/font.json")
}

func (m *MainScene) Setup(g *game.Game) {
	// set font
	gui.SetFont(assets.Font.GetFont("asc"))

	// image
	face := assets.Texture.Get("assets/face.png")
	m.face = face

	// image button background
	m.pressed = assets.Texture.Get("assets/particle.png")
	m.normal = assets.Texture.Get("assets/block.png")

	// slide default value
	m.slide = .5

	// input
	input.RegisterButton("A", input.A)
	input.RegisterButton("B", input.B)
}

func (m *MainScene) Update(dt float32) {
	m.NormalLayout()
}

func (m *MainScene) NormalLayout() {
	gui.Move(100, 50)
	gui.BeginHorizontal(1)

	// 针对当前 Group 的设置
	gui.SetGravity(.5, .5)
	gui.SetPadding(0,0,0,0)

	// gui.Cursor().set
	gui.Text(2, "Horizontal", nil)
	gui.BeginVertical(3)
	gui.SetGravity(m.slide, .5)
	gui.Text(4, "Vertical", nil)
	gui.Text(5, "Layout", nil)

	gui.Cursor().SetSize(30, 30).To(8)
	gui.Image(8, m.face, nil)

	if e := gui.Button(9, "NewButton", nil); (e & gui.EventWentDown) != 0 {
		log.Println("Click New Button")
		m.showbutton = true
	}
	if m.showbutton {
		if e := gui.Button(10, "Dismiss", nil); (e & gui.EventWentDown) != 0 {
			log.Println("Click Old Button")
			m.showbutton = false
		}
	}

	// image button
	gui.Cursor().SetSize(30, 30).To(12)
	gui.ImageButton(12, m.normal, m.pressed, nil)

	gui.Cursor().SetMargin(10, 0, 0, 0).SetSize(120, 9).To(11)
	gui.Slider(11, &m.slide, nil)

	gui.EndVertical()

	gui.Text(6, "Layout", nil)
	gui.EndHorizontal()

	// gui.DefaultContext().Layout.Dump()
}
var b bool


func (m *MainScene) Name() string {
	return "main"
}


func main() {
	fmt.Println("Hello World!!")
	log.Println("")

	options := &korok.Options{
		Title:"UI Test!",
		Width:480,
		Height:320,
	}
	korok.RunScene(options, &MainScene{})
}



