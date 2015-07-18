package main

func main() {
	g := NewGame()
	for g.Running {
		g.Update()
		g.Draw()
	}
	g.Destroy()
}
