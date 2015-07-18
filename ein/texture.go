package ein

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type Texture struct {
	width, height int32
	Dest          *sdl.Rect
	texture       *sdl.Texture
	Angle         float64
}

// TODO - Test this
// Animation is needed also!

func (g *Engine) NewTexture(filename string, x, y int32) *Texture {
	// todo infer size from filename if needed
	// then animiations could work - see how gosu does this
	surface, err := img.Load(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create surface: %s\n", err)
		os.Exit(1)
	}
	// TODO - abstract loading away from init
	// introduce manager class and cache texture loading
	// better abtracted and also performance
	texture, err := img.LoadTexture(g.Renderer, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(1)
	}

	rect := &sdl.Rect{x, y, surface.W, surface.H}

	return &Texture{
		width:   surface.W,
		height:  surface.H,
		texture: texture,
		Dest:    rect,
	}
}

func (t *Texture) Position() (int32, int32) {
	return t.Dest.X, t.Dest.Y
}

func (t *Texture) Size() (int32, int32) {
	return t.Dest.W, t.Dest.H
}

func (t *Texture) Draw(renderer *sdl.Renderer) {
	//renderer.Copy(t.texture, nil, t.Dest)
	// NB t.Dest is the frame from the entire texture
	renderer.CopyEx(t.texture, nil, t.Dest, t.Angle, nil, sdl.FLIP_NONE)
}
