package ein

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

// Drawable represents something that can be drawn, and placed in a Level.
type Drawable interface {
	//Tick(Event)         // Method for processing events, e.g. input
	Draw(*sdl.Renderer) // Method for drawing to the screen
}

// Physical represents something that can collide with another
// Physical, but cannot process its own collisions.
// Optional addition to Drawable.
type Physical interface {
	Position() (int32, int32) // Return position, x and y
	Size() (int32, int32)     // Return width and height
}

// DynamicPhysical represents something that can process its own collisions.
// Implementing this is an optional addition to Drawable.
type DynamicPhysical interface {
	Position() (int32, int32) // Return position, x and y
	Size() (int32, int32)     // Return width and height
	Collide(Physical)         // Handle collisions with another Physical
}

// Provides an event, for input
type Event struct {
	KeyCode *sdl.Keycode
}

type Engine struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
	Running       bool
	width, height int32
	Event         *sdl.Keycode

	entities []Drawable
}

func Init(width, height int32) *Engine {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int(width), int(height), sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create window", sdl.GetError())
		os.Exit(1)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	img.Init(0)
	ttf.Init()

	return &Engine{
		Window:   window,
		Renderer: renderer,
		Running:  true,
		width:    width,
		height:   height,
		entities: make([]Drawable, 0),
	}
}

func (g *Engine) AddEntities(d ...Drawable) {
	for _, e := range d {
		g.entities = append(g.entities, e)
	}
}

func (g *Engine) eventLoop() {
	g.Event = nil
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			g.Running = false

		case *sdl.KeyDownEvent:
			g.Event = &t.Keysym.Sym
		}
	}
}

func (g *Engine) Update() {
	g.eventLoop()

	// Handle collisions
	colls := make([]Physical, 0)
	dynamics := make([]DynamicPhysical, 0)

	for _, e := range g.entities {
		if p, ok := interface{}(e).(Physical); ok {
			colls = append(colls, p)
		}
		if p, ok := interface{}(e).(DynamicPhysical); ok {
			dynamics = append(dynamics, p)
		}
	}

	jobs := make(chan DynamicPhysical, len(dynamics))
	results := make(chan int, len(dynamics))
	for w := 0; w <= len(dynamics)/3; w++ {
		go checkCollisionsWorker(colls, jobs, results)
	}
	for _, p := range dynamics {
		jobs <- p
	}
	close(jobs)
	for r := 0; r < len(dynamics); r++ {
		<-results
	}
}

func (g *Engine) Draw() {
	for _, e := range g.entities {
		e.Draw(g.Renderer)
	}
}

func (g *Engine) Destroy() {
	g.Renderer.Destroy()
	g.Window.Destroy()
}

func checkCollisionsWorker(ps []Physical, jobs <-chan DynamicPhysical, results chan<- int) {
	for p := range jobs {
		for _, c := range ps {
			if c == p {
				continue
			}
			px, py := p.Position()
			cx, cy := c.Position()
			pw, ph := p.Size()
			cw, ch := c.Size()
			if px < cx+cw && px+pw > cx &&
				py < cy+ch && py+ph > cy {
				p.Collide(c)
			}
		}
		results <- 1
	}
}
