package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/kori-irrlicht/mana-engine/scene"
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
	sdl.Do(func() {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				game.running = false
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\tscan:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Keysym.Scancode)
			}
		}
	})

}

func (s *MainMenuScene) Update() {}

func (s *MainMenuScene) Render(float32) {
	inf, _ := fh.Get("font")
	font := inf.(*ttf.Font)
	sdl.Do(func() {
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
	})
}

func (s *MainMenuScene) Ready() bool {
	return true
}

var _ scene.Scene = &MainMenuScene{}
