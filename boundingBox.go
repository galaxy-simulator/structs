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

// BoundingBox is a struct defining the spatial outreach of a box
type BoundingBox struct {
	Center Vec2    // Center of the box
	Width  float64 // Width of the box
}

// NewBoundingBox returns a new Bounding Box using the centerpoint and the width given by the function parameters
func NewBoundingBox(center Vec2, width float64) BoundingBox {
	return BoundingBox{Center: center, Width: width}
}
