package diagonalnormalbounce

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raymath"
)

// DiagonalNormalBounce - Displays a vertical bouncing square
type DiagonalNormalBounce struct {
	currentPosition rl.Vector2
	startPosition   rl.Vector2
	endPosition     rl.Vector2
	velocity        rl.Vector2
	speed           float32
}

// NewDiagonalNormalBounce - return a new vertical bounce
func NewDiagonalNormalBounce() (b DiagonalNormalBounce) {
	b.startPosition = rl.Vector2{X: 50, Y: 110}
	b.endPosition = rl.Vector2{X: 140, Y: 200}
	b.currentPosition = rl.Vector2{X: 50, Y: 110}
	b.velocity = rl.Vector2{X: 1, Y: 1}
	raymath.Vector2Normalize(&b.velocity)
	b.speed = float32(10.0)
	return
}

// Update - calculate movement
func (b *DiagonalNormalBounce) Update() {
	// Move object
	b.currentPosition.X += float32(b.velocity.X) * float32(rl.GetFrameTime()) * b.speed
	b.currentPosition.Y += float32(b.velocity.Y) * float32(rl.GetFrameTime()) * b.speed

	if b.currentPosition.X >= b.endPosition.X && b.currentPosition.Y >= b.endPosition.Y {
		b.velocity.X *= -1
		b.velocity.Y *= -1
	}
	if b.currentPosition.X < b.startPosition.X || b.currentPosition.Y < b.startPosition.Y {
		b.velocity.X *= -1
		b.velocity.Y *= -1
	}

}

// Draw - render
func (b *DiagonalNormalBounce) Draw() {
	// Start
	rl.DrawRectangle(
		int32(b.startPosition.X), int32(b.startPosition.Y),
		int32(2), int32(2),
		rl.Red)
	// End
	rl.DrawRectangle(
		int32(b.endPosition.X), int32(b.endPosition.Y),
		int32(2), int32(2),
		rl.Green)

	// Object to move
	rl.DrawRectangle(
		int32(b.currentPosition.X-5), int32(b.currentPosition.Y-5),
		int32(10), int32(10),
		rl.Blue)
}

// SpeedUp - increase speed
func (b *DiagonalNormalBounce) SpeedUp() {
	b.speed *= 2
}

// SpeedDown - descrease speed
func (b *DiagonalNormalBounce) SpeedDown() {
	b.speed /= 2
}
