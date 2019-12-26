package angle

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Angle - structure for calculating angle between 2 x,y coords
type Angle struct {
	positionX     float64
	positionY     float64
	startPosition rl.Vector2
	endPosition   rl.Vector2
	drawPosition  rl.Vector2
	angle         float64
	speed         float64
	flipped       bool
}

// NewAngle - returns a new angle obj
func NewAngle(drawPosition rl.Vector2) Angle {
	a := Angle{}
	a.drawPosition = drawPosition
	a.startPosition = rl.Vector2{X: 0, Y: 200}
	a.endPosition = rl.Vector2{X: 200, Y: 0}
	a.positionX = float64(a.startPosition.X + 10)
	a.positionY = float64(a.startPosition.Y - 10)

	// Calculate angle
	a.angle = calculateAngle(a.startPosition, a.endPosition)
	fmt.Printf("Angle: %f\n", a.angle)

	a.speed = 1

	return a
}

// Draw - draw the stuffs
func (a *Angle) Draw() {
	// Draw starting position
	rl.DrawRectangle(
		int32(a.drawPosition.X+a.startPosition.X),
		int32(a.drawPosition.Y+a.startPosition.Y),
		1, 1, rl.Red)

	// Draw ending position
	rl.DrawRectangle(
		int32(a.drawPosition.X+a.endPosition.X),
		int32(a.drawPosition.Y+a.endPosition.Y),
		1, 1, rl.Green)

	// Draw ball
	rl.DrawRectangle(
		int32(a.drawPosition.X+float32(a.positionX)),
		int32(a.drawPosition.Y+float32(a.positionY)),
		10, 10, rl.Blue)

}

// Update - update the angle
func (a *Angle) Update() {
	a.positionX = math.Round(a.positionX + math.Cos(a.angle)*a.speed)
	a.positionY = math.Round(a.positionY - math.Sin(a.angle)*a.speed)
	fmt.Printf("X, Y: %f, %f\n", a.positionX, a.positionY)

	if float32(a.positionX) >= a.endPosition.X && float32(a.positionY) >= a.endPosition.Y {
		/*a.angle += 180
		a.flipped = true
		fmt.Printf("Inverting Angle: %f\n", a.angle)*/
		a.speed = 0
	}
	/*
		if float32(a.positionX) < a.startPosition.X || float32(a.positionY) < a.startPosition.Y {
			a.angle *= -1
			fmt.Printf("Inverting Angle: %f\n", a.angle)
		}
	*/
}

func calculateAngle(start, end rl.Vector2) float64 {
	var angle float64
	var theta float64

	theta = math.Atan2(float64(end.Y-start.Y), float64(end.X-start.X))
	fmt.Printf("Theta: %f\n", theta)
	theta += math.Pi / 2.0
	fmt.Printf("Theta + half Pi: %f\n", theta)

	angle = theta / (math.Pi / 180)
	if angle < 360 {
		angle += 360
	}
	if angle > 360 {
		angle -= 360

	}

	fmt.Printf("Angle in degrees: %f\n", angle)
	return angle
}
