package main

import "github.com/stephenbinns/ein/ein"

const velocity int32 = 5

type Bat struct {
	*ein.Rect
	score int
	YVel  int32
}

func NewBat(x int32) *Bat {
	rect := ein.NewRect(x, ScreenCenterY-50, 10, 100)
	return &Bat{Rect: rect}
}

func (b *Bat) MoveUp() {
	b.YVel = velocity
}

func (b *Bat) MoveDown() {
	b.YVel = -velocity
}

func (b *Bat) Update(g *Game) {
	if b.Rect.Y+b.Rect.H >= winHeight && b.YVel > 0 || b.Rect.Y <= 0 && b.YVel < 0 {
		b.YVel = 0
	}

	b.Rect.Y += b.YVel
}
