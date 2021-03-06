package loader

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/kori-irrlicht/mana-engine/asset"
	"github.com/veandco/go-sdl2/ttf"
)

type TrueTypeFont struct{}

func (l *TrueTypeFont) Load(name, file string, args map[string]string) (font interface{}, err error) {
	logrus.WithFields(logrus.Fields{
		"file": file,
		"name": name,
		"args": args,
	}).Debugln("Loading font")
	size, err := strconv.Atoi(args["size"])
	if err != nil {
		return
	}
	f, err := ttf.OpenFont(file, size)
	if err != nil {
		return
	}

	return f, err
}

// Enforce interface implementation
var (
	_ asset.Loader = (*TrueTypeFont)(nil)
)
