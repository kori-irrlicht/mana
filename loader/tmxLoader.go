package loader

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/fardog/tmx"
	"github.com/kori-irrlicht/mana-engine/asset"
)

type Tmx struct{}

func (l *Tmx) Load(name, file string, args map[string]string) (m interface{}, err error) {

	logrus.WithFields(logrus.Fields{
		"file": file,
		"name": name,
		"args": args,
	}).Debugln("Loading tmx file")

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	m, err = tmx.Decode(f)

	return m, err
}

// Enforce interface implementation
var (
	_ asset.Loader = (*Tmx)(nil)
)
