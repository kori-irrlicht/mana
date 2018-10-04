package main

import (
	"github.com/sirupsen/logrus"
	"github.com/kori-irrlicht/mana-engine/scene"
	"github.com/kori-irrlicht/mana/controller"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const NameMainMenu = "mainmenu"

type MainMenuScene struct{}

func (s *MainMenuScene) Entry() {
	_, err := fh.Load("font", "assets/fonts/DroidSans.ttf", map[string]string{"size": "32"})
	if err != nil {
		logrus.WithError(err).Panicln("Couldn't load font")
	}
}

func (s *MainMenuScene) Exit() {}

func (s *MainMenuScene) Input() {
	if game.controller.IsDown(controller.EXIT) {
		game.running = false
	}
	if game.controller.IsDown(controller.UP) {
		game.manager.Next(NameIngame)
	}

}

func (s *MainMenuScene) Update() {}

func (s *MainMenuScene) Render(float32) {
	inf, _ := fh.Get("font")
	font := inf.(*ttf.Font)
	surf, err := font.RenderUTF8_Blended("Hallo Welt", sdl.Color{255, 0, 0, 255})
	if err != nil {
		logrus.WithError(err).Errorln("Could not render utf8")
	}

	winSurf, err := game.window.GetSurface()
	if err != nil {
		logrus.WithError(err).Panicln("Could not get surface from window")
	}
	if err := surf.Blit(nil, winSurf, nil); err != nil {
		logrus.WithError(err).Panicln("Could not blit surface")
	}
	game.window.UpdateSurface()
}

func (s *MainMenuScene) Ready() bool {
	return true
}

var _ scene.Scene = &MainMenuScene{}
