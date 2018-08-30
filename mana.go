// Copyright 2018 kori-irrlicht
package main

import (
	"github.com/Sirupsen/logrus"
	mana "github.com/kori-irrlicht/mana-engine"
	"github.com/kori-irrlicht/mana-engine/asset"
	"github.com/kori-irrlicht/mana-engine/input"
	"github.com/kori-irrlicht/mana-engine/scene"
	"github.com/kori-irrlicht/mana/controller"
	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var fh = asset.NewHolder(&asset.TrueTypeFontLoader{})
var game *Game

type Game struct {
	manager    scene.Manager
	window     *sdl.Window
	running    bool
	controller input.Controller
}

func (g *Game) Input() {
	g.controller.Update()
	g.manager.Input()
}

func (g *Game) Update() {
	g.manager.Update()
}
func (g *Game) Render(delta float32) {
	g.manager.Render(delta)
}
func (g *Game) Running() bool {
	return g.running
}

func main() {
	initViper()
	initLogger()
	window := initSDL()

	defer sdl.Quit()
	defer ttf.Quit()
	defer window.Destroy()

	game = newGame(window)

	mana.Run(game)
}

func newGame(window *sdl.Window) *Game {
	g := &Game{}
	g.window = window
	g.running = true
	g.manager = scene.DefaultManager
	g.manager.Register(NameMainMenu, &MainMenuScene{})
	g.manager.StartWith(NameMainMenu)
	g.controller = controller.NewSdlKeyboard()
	return g
}

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.mana")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Panicln("Failed to read config file")
	}
}

func initSDL() *sdl.Window {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		logrus.WithError(err).Panicln("Failed to init sdl")
	}

	if err := ttf.Init(); err != nil {
		logrus.WithError(err).Panicln("Failed to init ttf")
	}

	window, err := sdl.CreateWindow("Mana", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, viper.GetInt("game.width"), viper.GetInt("game.height"), sdl.WINDOW_SHOWN)
	if err != nil {
		logrus.WithError(err).Panicln("Couldn't create window")
	}
	return window
}

func initLogger() {
	switch viper.GetString("debug.logging") {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.WarnLevel)
	}
}
