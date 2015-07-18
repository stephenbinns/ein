package main

import (
	"github.com/stephenbinns/ein/ein"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int32 = 800, 600
const ScreenCenterX, ScreenCenterY int32 = winWidth / 2, winHeight / 2

type Game struct {
	*ein.Engine
	bat1, bat2           *Bat
	ball                 *Ball
	bat1Label, bat2Label *ein.Label
}

func NewGame() *Game {
	e := ein.Init(winWidth, winHeight)

	divider := ein.NewRect((ScreenCenterX)-5, 0, 10, winHeight)
	font := ein.LoadFont("Courier New.ttf", 64)
	bat1Label := ein.NewLabel("0", 100, 50, font)
	bat2Label := ein.NewLabel("0", 600, 50, font)

	bat1 := NewBat(10)
	bat2 := NewBat(winWidth - 20)
	ball := NewBall()

	e.AddEntities(ball, bat1, bat2, bat1Label, bat2Label, divider)

	return &Game{
		Engine:    e,
		bat1:      bat1,
		bat2:      bat2,
		ball:      ball,
		bat1Label: bat1Label,
		bat2Label: bat2Label,
	}
}

func (g *Game) Update() {
	g.Engine.Update()
	if g.Event != nil {
		switch *g.Event {
		case sdl.K_a:
			g.bat1.MoveUp()
		case sdl.K_q:
			g.bat1.MoveDown()
		case sdl.K_l:
			g.bat2.MoveUp()
		case sdl.K_o:
			g.bat2.MoveDown()
		case sdl.K_ESCAPE:
			g.Engine.Running = false
		}
	}

	g.bat1.Update(g)
	g.bat2.Update(g)
	g.ball.Update(g)
}

func (g *Game) Draw() {
	g.Renderer.SetDrawColor(0, 0, 0, 255)
	g.Renderer.Clear()
	g.Engine.Draw()
	g.Renderer.Present()
}
