package structs

// Definition of a quadtree and it's nodes recursively
type Quadtree struct {
	boundary     BoundingBox // Spatial outreach of the quadtree
	centerOfMass Vec2        // Center of mass of the cell
	totalMass    float64     // Total mass of the cell
	depth        int         // Depth of the cell in the quadtree
	star         Star2D      // Star inside the cell

	// NW, NE, SW, SE
	quadrants []*Quadtree // List of quadtrees representing individual quadrants

	// Quadrants
	//northWest *Quadtree
	//northEast *Quadtree
	//southWest *Quadtree
	//southEast *Quadtree
}

// CenterOfMass is a getter method for quadtrees.
// It returns the Center of mass of the quadtree it is applied on
func (q *Quadtree) CenterOfMass() Vec2 {
	return q.centerOfMass
}

// SetCenterOfMass is a setter method for quadtrees.
// It sets the centerOfMass of the quadtree to the given value
func (q *Quadtree) SetCenterOfMass(centerOfMass Vec2) {
	q.centerOfMass = centerOfMass
}

// CalcCenterOfMass is a calculator method for quadtrees.
// It recursively walks through the quadtree and calculates it's center of mass.
// The calculated center of mass is then inserted into the centerOfMass variable.
func (q *Quadtree) CalcCenterOfMass() (Vec2, float64) {
	var totalMass float64 = 0
	var x float64 = 0
	var y float64 = 0

	// If the Node is a leaf
	if q.IsLeaf() == true {

		// update the values needed to calculate the center of mass
		totalMass += q.star.M
		x += q.star.C.X * q.star.M
		y += q.star.C.X * q.star.M

		return Vec2{x, y}, totalMass

	} else {

		// Iterate over all the quadrants
		for _, element := range q.quadrants {

			// Calculate the center of mass for each quadrant
			centerOfMass, totalMass := element.CalcCenterOfMass()

			// Update the overall centerOfMass for the individual quadtree
			q.centerOfMass.X += centerOfMass.X
			q.centerOfMass.Y += centerOfMass.Y
			q.totalMass += totalMass
		}
	}

	// Return the original centerOfMass and totalMass
	return q.centerOfMass, q.totalMass
}

// IsLeaf is a method for quadtrees returning true if the node is a leaf (has no children)
// or returning false if the node is nor a leaf (has children).
func (q *Quadtree) IsLeaf() bool {
	for _, element := range q.quadrants {
		if element == nil {
			return true
		}
	}
	return false
}

// NewQuadtree generates a new root node.
//func NewQuadtree(boundary BoundingBox) *Quadtree {
//	return &Quadtree{boundary: boundary}
//}
