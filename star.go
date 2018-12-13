package structs

// Define a struct storing essential star information such as it's coordinate, velocity and mass
type Star2D struct {
	C Vec2    // coordinates of the star
	V Vec2    // velocity    of the star
	M float64 // mass        of the star
}

// InsideOf is a method that tests if the star it is applied on is in or outside of the given
// BoundingBox. It returns true if the star is inside of the BoundingBox and false if it isn't.
func (s Star2D) InsideOf(boundary BoundingBox) bool {

	// Test if the star is inside or outside of the bounding box.
	// Abort testing if one of the conditions is not met
	if s.C.X < boundary.center.X+boundary.width/2 {
		if s.C.X > boundary.center.X-boundary.width/2 {
			if s.C.Y < boundary.center.Y+boundary.width/2 {
				if s.C.Y > boundary.center.Y-boundary.width/2 {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

// Quadrant returns a string indicating in which quadrant of the given quadtree the point the method
// is applied on is.
// This methods presumes that the point is inside of the boundingBox
func (s Star2D) Quadrant(starsQuadtree *Quadtree) string {
	centerX := starsQuadtree.boundary.center.X
	centerY := starsQuadtree.boundary.center.Y

	// test if the point is left the the center or not
	if s.C.X < centerX {

		// Test if the point is above or below of the center
		if s.C.Y > centerY {
			return "northwest"
		} else {
			return "southwest"
		}

		// The point is right of the center
	} else {

		// Test if the point is above or below of the center
		if s.C.Y > centerY {
			return "northeast"
		} else {
			return "southeast"
		}
	}
}

// Return a copy of the star by returning a star struct with the same values.
func (s *Star2D) Copy() Star2D {
	return Star2D{s.C.Copy(), s.V.Copy(), s.M}
}

// Accelerate the star with the acceleration a for the time t.
// This changes the velocity of the star.
func (s *Star2D) AccelerateVelocity(a Vec2, t float64) {
	s.V = s.V.Add(a.Multiply(t))
}

// Move the star with it's velocity for the time t.
// This changes the Position of the star.
func (s *Star2D) Move(t float64) {
	s.C = s.C.Add(s.V.Multiply(t))
}

// Accelerate and move the star with it's velocity and the acceleration a for the time t
// This changes the position and the velocity of the star.
func (s *Star2D) Accelerate(a Vec2, t float64) {
	s.AccelerateVelocity(a, t)
	s.Move(t)
}
