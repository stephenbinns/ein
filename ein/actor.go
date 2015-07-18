package ein

// import (
// 	"github.com/veandco/go-sdl2/sdl"
// )

// type Actor struct {
// 	Texture    *Texture
// 	XVel, YVel int32
// }

// func (g *Game) NewSingleFrameActor(filename string, x, y int32) *Actor {
// 	texture := g.NewTexture(filename, x, y)
// 	return NewActorFromTexture(texture)
// }

// func NewActorFromTexture(texture *Texture) *Actor {
// 	return &Actor{
// 		Texture: texture,
// 	}
// }

// func (b *Actor) Update() {
// 	b.Texture.Dest.X += b.XVel
// 	b.Texture.Dest.Y += b.YVel
// }

// func (b *Actor) Draw(renderer *sdl.Renderer) {
// 	b.Texture.Draw(renderer)
// }
