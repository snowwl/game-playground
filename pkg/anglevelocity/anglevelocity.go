package anglevelocity

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raymath"
)

// AngleVelocity - structure for calculating angle between 2 x,y coords
type AngleVelocity struct {
	position      rl.Vector2
	startPosition rl.Vector2
	endPosition   rl.Vector2
	drawPosition  rl.Vector2
	angle         float64
	speed         float32
	flipped       bool
	velocity      rl.Vector2
}

// NewAngleVelocity - returns a new angle obj
func NewAngleVelocity(drawPosition rl.Vector2) AngleVelocity {
	a := AngleVelocity{}
	a.drawPosition = drawPosition
	a.startPosition = rl.Vector2{X: 0, Y: 200}
	a.endPosition = rl.Vector2{X: 200, Y: 0}
	a.position.X = float32(a.startPosition.X + 10)
	a.position.Y = float32(a.startPosition.Y - 10)

	// Calculate angle
	a.angle = calculateAngle(a.startPosition, a.endPosition)
	fmt.Printf("AngleVelocity: %f\n", a.angle)

	a.velocity = rl.Vector2{
		X: float32(math.Cos(a.angle)),
		Y: float32(math.Sin(a.angle))}

	raymath.Vector2Normalize(&a.velocity)

	a.speed = 1

	return a
}

// Draw - draw the stuffs
func (a *AngleVelocity) Draw() {
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
		int32(a.drawPosition.X+float32(a.position.X)),
		int32(a.drawPosition.Y+float32(a.position.Y)),
		10, 10, rl.Blue)

}

// Update - update the angle
func (a *AngleVelocity) Update() {
	a.position.X += a.velocity.X * a.speed
	a.position.Y -= a.velocity.Y * a.speed
	if float32(a.position.X) >= a.endPosition.X && float32(a.position.Y) >= a.endPosition.Y {
		/*a.angle += 180
		a.flipped = true
		fmt.Printf("Inverting AngleVelocity: %f\n", a.angle)*/
		a.speed = 0
	}
	/*
		if float32(a.position.X) < a.startPosition.X || float32(a.position.Y) < a.startPosition.Y {
			a.angle *= -1
			fmt.Printf("Inverting AngleVelocity: %f\n", a.angle)
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

	fmt.Printf("AngleVelocity in degrees: %f\n", angle)
	return angle
}
