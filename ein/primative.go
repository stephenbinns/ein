package ein

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Rect struct {
	*sdl.Rect
	r, g, b, a uint8
}

func NewRect(x, y, w, h int32) *Rect {
	rect := &Rect{Rect: &sdl.Rect{x, y, w, h}}

	rect.SetColor(255, 255, 255, 255)
	return rect
}

func (t *Rect) SetColor(r, g, b, a uint8) {
	t.r = r
	t.g = g
	t.b = b
	t.a = a
}

func (t *Rect) Position() (int32, int32) {
	return t.Rect.X, t.Rect.Y
}

func (t *Rect) Size() (int32, int32) {
	return t.Rect.W, t.Rect.H
}

func (t *Rect) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(t.r, t.g, t.b, t.a)
	renderer.FillRect(t.Rect)
	renderer.DrawRect(t.Rect)
}
