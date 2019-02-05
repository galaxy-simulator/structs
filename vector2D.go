package structs

type Vec2 struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

// NewVec2 returns a new Vec2 using the given coordinates
func NewVec2(x float64, y float64) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

// Copy creates a copy of the vector
func (v *Vec2) Copy() Vec2 {
	return Vec2{v.X, v.Y}
}

// returns the product of the vector and a scalar s
func (v *Vec2) Multiply(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// returns the sum of this vector and the vector v2
func (v1 *Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1.X + v2.X, v1.Y + v2.Y}
}
