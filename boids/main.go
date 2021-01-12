package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount = 500
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	boids[boidCount] * Boid
)

func update(screen *ebiten.Image) error {

	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
			screen.Set(int(boid.position.x), int(boid.position.y-1), green)
			screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		}
	}

	return nil

}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	if err := ebiten.Run(update, screenHeight, screenHeight, 2, "Boids in a box"); err != nil {
		log.Fatal(err)
	}
}
