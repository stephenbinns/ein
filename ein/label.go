package ein

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type Label struct {
	text string
	rect *sdl.Rect
	font *ttf.Font
}

func NewLabel(text string, x, y int32, font *ttf.Font) *Label {
	// TODO is is nice to pass fonts about like this
	// introduce font manager similar to texture and simply
	// take the font name
	return &Label{text, &sdl.Rect{x, y, 0, 0}, font}
}

func LoadFont(name string, size int) *ttf.Font {
	font, err := ttf.OpenFont(name, size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load font: %s\n", sdl.GetError())
		os.Exit(3)
	}

	return font
}

func (t *Label) SetText(text string) {
	t.text = text
}

func (t *Label) Draw(renderer *sdl.Renderer) {
	surface, err := t.font.RenderUTF8_Solid(t.text, sdl.Color{255, 255, 255, 255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create surface: %s\n", err)
		os.Exit(1)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(1)
	}

	src := &sdl.Rect{0, 0, surface.W, surface.H}
	dst := &sdl.Rect{t.rect.X, t.rect.Y, src.W, src.H}

	renderer.Copy(texture, src, dst)
	surface.Free()
}

func (t *Label) Position() (int32, int32) {
	return t.rect.X, t.rect.Y
}

func (t *Label) Size() (int32, int32) {
	return t.rect.W, t.rect.H
}
