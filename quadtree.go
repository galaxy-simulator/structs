package structs

// Definition of a quadtree and it's nodes recursively
type Quadtree struct {
	Boundary     BoundingBox `json:"boundary"`     // Spatial outreach of the quadtree
	CenterOfMass Vec2        `json:"CenterOfMass"` // Center of mass of the cell
	TotalMass    float64     `json:"totalMass"`    // Total mass of the cell
	Depth        int         `json:"depth"`        // Depth of the cell in the quadtree
	Star         Star2D      `json:"star"`         // Star inside the cell

	// NW, NE, SW, SE
	Quadrants []*Quadtree `json:"Quadrants"` // List of quadtrees representing individual Quadrants

	// Quadrants
	//northWest *Quadtree
	//northEast *Quadtree
	//southWest *Quadtree
	//southEast *Quadtree
}

// SetCenterOfMass is a setter method for quadtrees.
// It sets the CenterOfMass of the quadtree to the given value
func (q *Quadtree) SetCenterOfMass(centerOfMass Vec2) {
	q.CenterOfMass = centerOfMass
}

// CalcCenterOfMass is a calculator method for quadtrees.
// It recursively walks through the quadtree and calculates it's Center of mass.
// The calculated Center of mass is then inserted into the CenterOfMass variable.
func (q *Quadtree) CalcCenterOfMass() (Vec2, float64) {
	var totalMass float64 = 0
	var x float64 = 0
	var y float64 = 0

	// If the Node is a leaf
	if q.IsLeaf() == true {

		// update the values needed to calculate the Center of mass
		totalMass += q.Star.M
		x += q.Star.C.X * q.Star.M
		y += q.Star.C.X * q.Star.M

		return Vec2{x, y}, totalMass

	} else {

		// Iterate over all the Quadrants
		for _, element := range q.Quadrants {

			// Calculate the Center of mass for each quadrant
			centerOfMass, totalMass := element.CalcCenterOfMass()

			// Update the overall CenterOfMass for the individual quadtree
			q.CenterOfMass.X += centerOfMass.X
			q.CenterOfMass.Y += centerOfMass.Y
			q.TotalMass += totalMass
		}
	}

	// Return the original CenterOfMass and totalMass
	return q.CenterOfMass, q.TotalMass
}

// IsLeaf is a method for quadtrees returning true if the node is a leaf (has no children)
// or returning false if the node is nor a leaf (has children).
func (q *Quadtree) IsLeaf() bool {
	for _, element := range q.Quadrants {
		if element == nil {
			return true
		}
	}
	return false
}

// NewQuadtree generates a new root node.
func NewQuadtree(boundary BoundingBox) *Quadtree {
	return &Quadtree{Boundary: boundary}
}
