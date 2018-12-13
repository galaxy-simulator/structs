package structs

import (
	"math"
)

type Vec2 struct {
	X, Y float64
}

// newVec2 returns a new Vec2 using the given coordinates
func newVec2(x float64, y float64) *Vec2 {
	return &Vec2{
		X: x,
		Y: y,
	}
}

// creates a copy of the vector
func (v *Vec2) Copy() Vec2 {
	return Vec2{v.X, v.Y}
}

func (v *Vec2) Split() (x float64, y float64) {
	return v.X, v.Y
}

// changes the length of the vector to the length l
func (v *Vec2) SetLength(l float64) {
	var k = l / v.GetLength()
	var newV = v.Multiply(k)
	//	v = newV
	v.X, v.Y = newV.Split()
}

// changes the length of the vector to the length 1
func (v *Vec2) SetLengthOne() {
	v.SetLength(1)
}

// returns the direction Vector of this vector. This means a copy of this vector with a length of 1
func (v *Vec2) GetDirVector() Vec2 {
	var dirV = v.Copy()
	dirV.SetLengthOne()
	return dirV
}

// returns the length of the vector
func (v *Vec2) GetLength() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// returns the product of the vector and a scalar s
func (v *Vec2) Multiply(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// returns the quotient of the vector and a scalar s
func (v *Vec2) Divide(s float64) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

// returns the sum of this vector and the vector v2
func (v1 *Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

// returns the difference of this vector minus the vector v2
func (v1 *Vec2) Subtract(v2 Vec2) Vec2 {
	return Vec2{v1.X - v2.X, v1.Y - v2.Y}
}
