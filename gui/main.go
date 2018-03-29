package main

import (
	"korok.io/korok/gui"
	"korok.io/korok/hid/input"
	"korok.io/korok/asset"
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
	asset.Texture.Load("assets/face.png")
	asset.Texture.Load("assets/block.png")
	asset.Texture.Load("assets/particle.png")
	asset.Font.LoadBitmap("asc", "assets/font/font.png", "assets/font/font.json")
}

func (m *MainScene) OnEnter(g *game.Game) {
	// set font
	gui.SetFont(asset.Font.GetFont("asc"))

	// image
	face := asset.Texture.Get("assets/face.png")
	m.face = face

	// image button background
	m.pressed = asset.Texture.Get("assets/particle.png")
	m.normal = asset.Texture.Get("assets/block.png")

	// slide default value
	m.slide = .5

	// input
	input.RegisterButton("A", input.A)
	input.RegisterButton("B", input.B)
}

func (m *MainScene) Update(dt float32) {
	//m.Widget()
	m.Layout()
}

func (m *MainScene) OnExit() {
	return
}



func (m *MainScene) Widget() {

	gui.Move(100, 60)
	gui.Layout(1, func(g *gui.Group, p *gui.Params) {
		gui.Text(2, "SomeText", nil)

		p.SetSize(30, 30).To(3)
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
		p.SetSize(30, 30).To(6)
		gui.ImageButton(6, m.normal, m.pressed, nil)

		p.SetSize(120, 9).To(7)
		gui.Slider(7, &m.slide, nil)
	}, 0, 0, gui.Vertical)

	// gui.DefaultContext().Layout.Dump()
}

// show how to layout ui-element
func (m *MainScene) Layout() {
	gui.Move(0, 0)
	gui.Layout(0, func(g *gui.Group, p *gui.Params) {
		p.SetGravity(.5, 0).To(1)
		gui.Text(1, "Top", nil)

		p.SetGravity(.5, 1).To(2)
		gui.Text(2, "Bottom", nil)

		p.SetGravity(0, .5).To(3)
		gui.Text(3, "Left", nil)

		p.SetGravity(1, .5).To(4)
		gui.Text(4, "Right", nil)

		p.SetGravity(.5, .5).To(5)

		gui.Layout(5, func(g *gui.Group, p *gui.Params) {
			gui.Text(6, "Horizontal", nil)

			gui.Layout(7, func(g *gui.Group, p *gui.Params) {
				gui.Text(8, "Vertical", nil)
				gui.Text(9, "Layout", nil)
			}, 0, 0, gui.Vertical)

			gui.Text(10, "Layout", nil)
		}, 0, 0, gui.Horizontal)
	}, 480 - 16, 320 - 16, gui.OverLay)
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



