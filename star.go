// star.go defines stars and actions that can be used on them
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

// Star2D defines a struct storing essential star information such as it's coordinate, velocity and mass
type Star2D struct {
	C Vec2    `json:"C"` // coordinates of the star
	V Vec2    `json:"V"` // velocity    of the star
	M float64 `json:"M"` // mass        of the star
}

// NewStar2D returns a new star using the given arguments as values for the Star
func NewStar2D(c Vec2, v Vec2, m float64) Star2D {
	return Star2D{C: c, V: v, M: m}
}

// InsideOf is a method that tests if the star it is applied on is in or outside of the given
// BoundingBox. It returns true if the star is inside of the BoundingBox and false if it isn't.
func (star Star2D) InsideOf(boundary BoundingBox) bool {

	// Test if the star is inside or outside of the bounding box.
	// Abort testing if one of the conditions is not met
	if star.C.X < boundary.Center.X+boundary.Width/2 {
		if star.C.X > boundary.Center.X-boundary.Width/2 {
			if star.C.Y < boundary.Center.Y+boundary.Width/2 {
				if star.C.Y > boundary.Center.Y-boundary.Width/2 {
					return true
				}
			}
		}
	}

	return false
}

// CalcNewPos calculates the new position of a star using the force acting on it
func (star *Star2D) CalcNewPos(force Vec2, timestep float64) {
	acceleration := NewVec2(force.X/star.M, force.Y/star.M)
	star.Accelerate(acceleration, timestep)
}

// Copy Return a copy of the star by returning a star struct with the same values.
func (star *Star2D) Copy() Star2D {
	return Star2D{star.C.Copy(), star.V.Copy(), star.M}
}

// AccelerateVelocity accelerates the star with the acceleration a for the time t.
// This changes the velocity of the star.
func (star *Star2D) AccelerateVelocity(a Vec2, t float64) {
	star.V = star.V.Add(a.Multiply(t))
}

// Move the star with it's velocity for the time t.
// This changes the Position of the star.
func (star *Star2D) Move(t float64) {
	star.C = star.C.Add(star.V.Multiply(t))
}

// Accelerate and move the star with it's velocity and the acceleration a for the time t
// This changes the position and the velocity of the star.
func (star *Star2D) Accelerate(a Vec2, t float64) {
	star.AccelerateVelocity(a, t)
	star.Move(t)
}

// posX determines if the star is the positive x region of the given boundary. If it is,
// the method returns true, if not, it returns false
func (star Star2D) posX(boundary BoundingBox) bool {

	// define shortcuts
	bx := boundary.Center.X
	bw := boundary.Width / 2

	if star.C.X > bx && star.C.X < bx+bw {
		return true
	}
	return false
}

// posY determines if the star is the positive y region of the given boundary. If it is,
// the method returns true, if not, it returns false
func (star Star2D) posY(boundary BoundingBox) bool {

	// define shortcuts
	by := boundary.Center.Y
	bw := boundary.Width / 2

	if star.C.Y > by && star.C.Y < by+bw {
		return true
	}
	return false
}

// getRelativePosition returns the relative position of a star relative to the bounding
// bounding box it is in. It returns the integer that is mapped to a cell in the Node
// definition
func (star Star2D) getRelativePosition(boundary BoundingBox) string {
	if star.posX(boundary) == true {
		if star.posY(boundary) == true {
			return "NE"
		}
		return "SE"
	}
	if star.posY(boundary) == true {
		return "NW"
	}
	return "SW"
}

func (star Star2D) getRelativePositionInt(boundary BoundingBox) int {
	quadrantMap := make(map[string]int)
	quadrantMap["NW"] = 0
	quadrantMap["NE"] = 1
	quadrantMap["SW"] = 2
	quadrantMap["SE"] = 3

	QuadrantMapString := star.getRelativePosition(boundary)
	return quadrantMap[QuadrantMapString]
}
