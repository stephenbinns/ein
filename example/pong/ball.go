package main

import (
	"math"
	"strconv"

	"github.com/stephenbinns/ein/ein"
)

const initialVelocity int32 = 5

type Ball struct {
	*ein.Rect
	xvel, yvel int32
}

func NewBall() *Ball {
	rect := ein.NewRect(ScreenCenterX, ScreenCenterY, 10, 10)
	return &Ball{rect, initialVelocity, initialVelocity}
}

func (b *Ball) Update(g *Game) {
	b.Rect.X += b.xvel
	b.Rect.Y += b.yvel

	if b.Rect.Y >= winHeight || b.Rect.Y <= 0 {
		b.yvel *= -1
	}

	if b.Rect.X >= winWidth {
		b.score(g.bat1, g.bat1Label)
	}

	if b.Rect.X <= 0 {
		b.score(g.bat2, g.bat2Label)
	}
}

func (b *Ball) score(bat *Bat, label *ein.Label) {
	b.Rect.X, b.Rect.Y = ScreenCenterX, ScreenCenterY
	b.xvel, b.yvel = initialVelocity, initialVelocity
	bat.score++
	label.SetText(strconv.Itoa(bat.score))
}

func (b *Ball) Collide(p ein.Physical) {
	// check we collided with a bat
	if a, ok := p.(*Bat); ok {
		b.xvel *= -1
		b.yvel *= -(a.YVel / 2)

		paddleHeight := a.Rect.H

		intersect := float64((a.Rect.Y + (paddleHeight / 2.0)) - b.Rect.Y)
		normalized := float64(intersect / float64(paddleHeight/2.0))
		bounceAngle := normalized * 65.0 // 65 deg max

		// if the bat is the opposite side rotate by 180
		if a.Rect.X > 500 {
			bounceAngle += 180
		}

		// convert to radians - like a boss.
		bounceAngle *= math.Pi / 180

		b.xvel = int32(8.0 * math.Cos(float64(bounceAngle)))
		b.yvel = int32(8.0 * -math.Sin(float64(bounceAngle)))
	}
}
