package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kori-irrlicht/mana-engine/scene"
	"github.com/kori-irrlicht/mana/controller"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const NameIngame = "ingame"

type IngameScene struct {
	font *ttf.Font
}

func (s *IngameScene) Entry() {
	_, err := fh.Load("font", "assets/fonts/DroidSans.ttf", map[string]string{"size": "32"})
	if err != nil {
		logrus.WithError(err).Errorln("Couldn't load font")
	}

	inf, _ := fh.Get("font")
	s.font = inf.(*ttf.Font)

}

func (s *IngameScene) Exit() {}

func (s *IngameScene) Input() {
	if game.controller.IsDown(controller.EXIT) {
		game.running = false
	}

}

func (s *IngameScene) Update() {}

func (s *IngameScene) Render(float32) {

	surf, err := s.font.RenderUTF8_Blended("Ingame", sdl.Color{255, 0, 0, 255})
	if err != nil {
		logrus.WithError(err).Errorln("Could not render utf8")
	}

	winSurf, err := game.window.GetSurface()
	if err != nil {
		logrus.WithError(err).Panicln("Could not get surface from window")
	}
	winSurf.FillRect(nil, 0xffffff)
	if err := surf.Blit(nil, winSurf, nil); err != nil {
		logrus.WithError(err).Panicln("Could not blit surface")
	}
	game.window.UpdateSurface()

}

func (s *IngameScene) Ready() bool {
	return true
}

var _ scene.Scene = (*IngameScene)(nil)
