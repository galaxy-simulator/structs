// main.go purpose is to build the interaction layer in between the http endpoints and the http server
// Copyright (C) 2019 Emile Hansmaennel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

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

// Multiply returns the product of the vector and a scalar s
func (v *Vec2) Multiply(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// Add returns the sum of this vector and the vector v2
func (v *Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}
