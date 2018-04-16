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
	"korok.io/korok/gui/auto"
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
	asset.Texture.Load("face.png")
	asset.Texture.Load("block.png")
	asset.Texture.Load("particle.png")
	asset.Font.LoadBitmap("font1", "font.png", "font.json")
}

func (m *MainScene) OnEnter(g *game.Game) {
	// set font
	gui.SetFont(asset.Font.GetFont("font1"))

	// image
	face := asset.Texture.Get("face.png")
	m.face = face

	// image button background
	m.pressed = asset.Texture.Get("particle.png")
	m.normal = asset.Texture.Get("block.png")

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

	gui.Move(100, 60)

	// draw text
	gui.Text(2, gui.Rect{0, 0, 0, 0}, "SomeText", nil)

	// draw image
	gui.Image(3, gui.Rect{0, 30, 30, 30}, m.face, nil)

	// draw image button
	gui.ImageButton(6, gui.Rect{50, 30, 30, 30}, m.normal, m.pressed, nil)


	// draw button
	if e := gui.Button(4, gui.Rect{0, 100, 0, 0},"NewButton", nil); (e & gui.EventWentDown) != 0 {
		log.Println("Click New Button")
		m.showbutton = true
	}
	if m.showbutton {
		if e := gui.Button(5, gui.Rect{0, 150, 0, 0},"Dismiss", nil); (e & gui.EventWentDown) != 0 {
			log.Println("Click Old Button")
			m.showbutton = false
		}
	}

	// draw slider
	gui.Slider(7, gui.Rect{100, 0, 120, 10},  &m.slide, nil)

	// gui.DefaultContext().Layout.Dump()
}

// show how to layout ui-element
func (m *MainScene) Layout() {
	gui.Move(0, 0)
	auto.Layout(0, func(g *auto.Group) {
		auto.Text(1, "Top", nil, auto.Gravity(.5, 0))
		auto.Text(2, "Bottom", nil, auto.Gravity(.5, 1))
		auto.Text(3, "Left", nil, auto.Gravity(0, .5))
		auto.Text(4, "Right", nil, auto.Gravity(1, .5))

		// p.SetGravity(.5, .5).To(5)

		auto.Layout(5, func(g *auto.Group) {
			auto.Text(6, "Horizontal", nil, nil)

			auto.Layout(7, func(g *auto.Group) {
				auto.Text(8, "Vertical", nil, nil)
				auto.Text(9, "Layout", nil, nil)
			}, 0, 0, auto.Vertical)

			auto.Text(10, "Layout", nil, nil)
		}, 0, 0, auto.Horizontal)
	}, 480 - 16, 320 - 16, auto.OverLay)
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



