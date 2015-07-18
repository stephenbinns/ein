package main

import (
	"github.com/stephenbinns/ein/ein"
	"github.com/veandco/go-sdl2/sdl"
)

type Bullet struct {
	*ein.Rect
	angle, speed int32
}

func (p *Bullet) Update() {

}

type Player struct {
	*ein.Texture
	score, speed int32
}

func NewPlayer(e *ein.Engine) *Player {
	texture := e.NewTexture("player.png", 400, 300)
	player := &Player{Texture: texture}
	e.AddEntities(texture)
	return player
}

func (p *Player) TurnLeft() {
	p.Angle += 5
}

func (p *Player) TurnRight() {
	p.Angle -= 5
}

func (p *Player) Booster() {

}

func (p *Player) Update() {

}

func (p *Player) Shoot(e *ein.Engine) {

}

func (p *Player) Collides(c ein.Physical) {
	/*if a, ok := c.(*Asteroid); ok {
	}*/
}

type Asteroid struct {
	*ein.Texture
}

func (p *Asteroid) Update() {

}

func (a *Asteroid) Collides(p ein.Physical) {
	/*if b, ok := p.(*Bullet); ok {
	}

	if other, ok := p.(*Asteroid); ok {
	}*/
}

type Game struct {
	*ein.Engine
	*Player
}

func NewGame() *Game {
	e := ein.Init(800, 600)
	player := NewPlayer(e)
	return &Game{
		Engine: e,
		Player: player,
	}
}

func (g *Game) Update() {
	g.Engine.Update()
	if g.Event != nil {
		switch *g.Event {
		case sdl.K_UP:
			g.Player.Booster()
		case sdl.K_LEFT:
			g.Player.TurnLeft()
		case sdl.K_RIGHT:
			g.Player.TurnRight()
		case sdl.K_SPACE:
			g.Player.Shoot(g.Engine)
		case sdl.K_ESCAPE:
			g.Engine.Running = false
		}
	}
}

func (g *Game) Draw() {
	g.Renderer.SetDrawColor(0, 0, 0, 255)
	g.Renderer.Clear()
	g.Engine.Draw()
	g.Renderer.Present()
}
