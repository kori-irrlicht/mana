package controller

import (
	"fmt"

	"github.com/kori-irrlicht/mana-engine/input"
	"github.com/veandco/go-sdl2/sdl"
)

type sdlKeyboard struct {
	keys map[input.Key]bool
}

func NewSdlKeyboard() input.Controller {
	sk := &sdlKeyboard{}
	sk.keys = make(map[input.Key]bool)
	return sk
}

func (sk sdlKeyboard) IsDown(key input.Key) bool {
	val, ok := sk.keys[key]
	if !ok {
		return false
	}
	return val
}

func (sk *sdlKeyboard) Update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			sk.keys[EXIT] = true
		case *sdl.KeyUpEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\tscan:%d\n",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Keysym.Scancode)
			sk.updateKeys(t.Keysym.Sym, false)
		case *sdl.KeyDownEvent:
			sk.updateKeys(t.Keysym.Sym, true)
		}
	}
}

func (sk *sdlKeyboard) updateKeys(sym sdl.Keycode, down bool) {
	switch sym {
	case sdl.K_w:
		sk.keys[UP] = down
	case sdl.K_a:
		sk.keys[LEFT] = down
	case sdl.K_s:
		sk.keys[DOWN] = down
	case sdl.K_d:
		sk.keys[RIGHT] = down
	case sdl.K_ESCAPE:
		sk.keys[EXIT] = down
	}
}

var _ input.Controller = (*sdlKeyboard)(nil)
