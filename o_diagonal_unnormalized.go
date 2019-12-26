package main

import rl "github.com/gen2brain/raylib-go/raylib"

// DiagonalBounce - Displays a vertical bouncing square
type DiagonalBounce struct {
	currentPosition rl.Vector2
	startPosition   rl.Vector2
	endPosition     rl.Vector2
	velocity        rl.Vector2
	speed           float32
}

// NewDiagonalBounce - return a new vertical bounce
func NewDiagonalBounce() (b DiagonalBounce) {
	b.startPosition = rl.Vector2{X: 10, Y: 110}
	b.endPosition = rl.Vector2{X: 100, Y: 200}
	b.currentPosition = rl.Vector2{X: 10, Y: 110}
	b.velocity = rl.Vector2{X: 1, Y: 1}
	b.speed = float32(10.0)
	return
}

// Update - calculate movement
func (b *DiagonalBounce) Update() {
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
func (b *DiagonalBounce) Draw() {
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
func (b *DiagonalBounce) SpeedUp() {
	b.speed *= 2
}

// SpeedDown - descrease speed
func (b *DiagonalBounce) SpeedDown() {
	b.speed /= 2
}
