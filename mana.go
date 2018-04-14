// Copyright 2018 kori-irrlicht
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	initViper()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		logrus.WithError(err).Panicln("Failed to init sdl")
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Mana", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, viper.GetInt("game.width"), viper.GetInt("game.height"), sdl.WINDOW_SHOWN)
	if err != nil {
		logrus.WithError(err).Panicln("Couldn't create window")
	}
	defer window.Destroy()

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
