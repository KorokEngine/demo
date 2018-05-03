package main

import (
	"korok.io/korok/game"
	"korok.io/korok"
	"korok.io/korok/asset"
	"korok.io/korok/gui"
	"korok.io/korok/audio"

	"log"
	"fmt"
)

func main() {
	fmt.Println("Hello Audio!!")
	options := korok.Options{
		Width:320,
		Height:480,
		Title:"Audio Test",
	}
	korok.Run(&options, &MainScene{})
}

type MainScene struct {
	wav uint16
	ogg uint16
}

func (*MainScene) Load() {
	asset.Font.LoadTrueType("ttf", "OCRAEXT.TTF")

	asset.Audio.Load("birds.wav", true)
	asset.Audio.Load("ambient.ogg", true)
}

func (m *MainScene) OnEnter(g *game.Game) {
	font, _ := asset.Font.Get("ttf")
	gui.SetFont(font)
	gui.SetVirtualResolution(320, 0)

	m.wav = asset.Audio.Get("birds.wav")
	m.ogg = asset.Audio.Get("ambient.ogg")
}

func (m *MainScene) Update(dt float32) {
	if gui.Button(1, gui.Rect{100, 100, 0, 0}, "Play", nil).JustPressed() {
		audio.PlayMusic(m.ogg)
		log.Println("play audio")
	}

	if gui.Button(2, gui.Rect{100, 140, 0, 0}, "Stop", nil).JustPressed() {
		// stop audio
		audio.StopMusic()
		log.Println("stop audio")
	}

	if gui.Button(3, gui.Rect{180, 100, 0, 0}, "Pause", nil).JustPressed() {
		audio.PauseMusic()
		log.Println("pause audio")
	}

	if gui.Button(4, gui.Rect{180, 140, 0, 0}, "Resume", nil).JustPressed() {
		audio.ResumeMusic()
		log.Println("resume audio")
	}
}

func (m *MainScene) OnExit() {
	audio.Destroy()
}




