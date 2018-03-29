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

// Note:
// To manager ui-element's id is hard in imgui. ref: https://gist.github.com/niklas-ourmachinery/9e37bdcad5bacaaf09ad4f5bb93ecfaf
// So I leave the problem to the developer.
// The id start from 1 (0 is used by the default layout).
type MainScene struct {
	face gfx.Tex2D
	slide float32

	normal, pressed gfx.Tex2D
	showbutton bool
}

func (m *MainScene) Load() {
	assets.Texture.Load("assets/face.png")
	assets.Texture.Load("assets/block.png")
	assets.Texture.Load("assets/particle.png")
	assets.Font.LoadBitmap("asc", "assets/font/font.png", "assets/font/font.json")
}

func (m *MainScene) OnEnter(g *game.Game) {
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
	m.Widget()
	//m.Layout()
}

func (m *MainScene) OnExit() {
	return
}



func (m *MainScene) Widget() {

	gui.Cursor().SetGravity(.5, .5).To(1)
	gui.BeginVertical(1)

	gui.Text(2, "SomeText", nil)

	gui.Cursor().SetSize(30, 30).To(3)
	gui.Image(3, m.face, nil)

	if e := gui.Button(4, "NewButton", nil); (e & gui.EventWentDown) != 0 {
		log.Println("Click New Button")
		m.showbutton = true
	}
	if m.showbutton {
		if e := gui.Button(5, "Dismiss", nil); (e & gui.EventWentDown) != 0 {
			log.Println("Click Old Button")
			m.showbutton = false
		}
	}

	// image button
	gui.Cursor().SetSize(30, 30).To(6)
	gui.ImageButton(6, m.normal, m.pressed, nil)

	gui.Cursor().SetSize(120, 9).To(7)
	gui.Slider(7, &m.slide, nil)

	gui.EndVertical()
	// gui.DefaultContext().Layout.Dump()
}

// show how to layout ui-element
func (m *MainScene) Layout() {
	gui.Cursor().SetGravity(.5, 0).To(1)
	gui.Text(1, "Top", nil)

	gui.Cursor().SetGravity(.5, 1).To(2)
	gui.Text(2, "Bottom", nil)

	gui.Cursor().SetGravity(0, .5).To(3)
	gui.Text(3, "Left", nil)

	gui.Cursor().SetGravity(1, .5).To(4)
	gui.Text(4, "Right", nil)

	gui.Cursor().SetGravity(.5, .5).To(5)
	gui.BeginHorizontal(5)
	gui.Text(6, "Horizontal", nil)

	gui.BeginVertical(7)
	gui.Text(8, "Vertical", nil)
	gui.Text(9, "Layout", nil)
	gui.EndVertical()

	gui.Text(10, "Layout", nil)
	gui.EndHorizontal()
}

var b bool

func main() {
	fmt.Println("Hello World!!")
	log.Println("")

	options := &korok.Options{
		Title:"UI Test!",
		Width:480,
		Height:320,
	}
	korok.Run(options, &MainScene{})
}



