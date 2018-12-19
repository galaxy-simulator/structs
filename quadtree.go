package structs

import (
	"log"
)

// Definition of a quadtree and it's nodes recursively
type Quadtree struct {
	Boundary     BoundingBox `json:"boundary"`     // Spatial outreach of the quadtree
	CenterOfMass Vec2        `json:"CenterOfMass"` // Center of mass of the cell
	TotalMass    float64     `json:"totalMass"`    // Total mass of the cell
	Depth        int         `json:"depth"`        // Depth of the cell in the quadtree
	Star         Star2D      `json:"star"`         // Star inside the cell
	Leaf         bool        `json:"Leaf"`         // Quadtree is a leaf or not

	// NW, NE, SW, SE
	Quadrants [4]*Quadtree `json:"Quadrants"` // List of quadtrees representing individual Quadrants

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

	q.IsLeaf()
	// If the Node is a leaf
	if q.Leaf == true {

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
func (q *Quadtree) IsLeaf() {

	// assume that the node is a leaf
	q.Leaf = true

	// iterate over all the elements in the quadtree (all the quadrants)
	for _, element := range q.Quadrants {

		// if one of the quadrants is not nil , the node is not a leaf
		if element != nil {
			q.Leaf = false
		}
	}
}

// NewQuadtree generates a new root node.
func NewQuadtree(boundary BoundingBox) *Quadtree {
	newquadtree := &Quadtree{
		Boundary: BoundingBox{
			Center: Vec2{
				X: boundary.Center.X,
				Y: boundary.Center.Y,
			},
			Width: boundary.Width,
		},
		CenterOfMass: Vec2{},
		TotalMass:    0,
		Depth:        0,
		Star:         Star2D{},
		Leaf:         true,
		Quadrants:    [4]*Quadtree{},
	}
	log.Printf("[+++] New Quadtree: %v", newquadtree)
	return newquadtree
}

// Insert inserts the given point into the quadtree the method is called on
func (q *Quadtree) Insert(point Star2D) {
	log.Printf("[   ] Inserting point %v into the tree %v", point, q)

	// prints the stars in the leaves of the current node
	log.Printf("[>>>] - Current Star [node]: %v", q.Star)
	for i := 0; i < 4; i++ {
		if q.Quadrants[i] != nil {
			log.Printf("[>>>] - Current Star [%d]: %v", i, q.Quadrants[i].Star)
		} else {
			log.Printf("[>>>] - Current Star [%d]: ", i)
		}
	}

	// create a shortcut for the various bounding box variables
	bx := q.Boundary.Center.X
	by := q.Boundary.Center.Y
	bw := q.Boundary.Width
	log.Printf("[ ~ ] \t Bounding Box X: %f", bx)
	log.Printf("[ ~ ] \t Bounding Box Y: %f", by)
	log.Printf("[ ~ ] \t Bounding Box Width: %f", bw)

	// Insert thee given star into the galaxy
	// Case 1: There is no star inside of the node
	if q.Star == (Star2D{Vec2{}, Vec2{}, 0}) {
		log.Printf("[ + ] There was no star inside of the node -> inserting directly")
		q.Star = point
		return

		// Case 2: There is all ready a star inside of the node
	} else {
		log.Printf("[ + ] There is allready a star inside of the node -> subdividing")

		// Test if the star is left or right of the center point
		if point.C.X < bx && point.C.X > bx-bw {
			log.Println("[<  ] \t\t The point is left of the center point!")

			// Test if the star is above or below the center point
			if point.C.Y > by && point.C.Y < by+bw {
				log.Println("[ ^ ] \t\t The point is above of the center point!")

				// Subdivide if...
				// ... the quadrant does not contain a node yet or if there is all ready a node in the quadrant
				// ... the quadrant is a leaf
				if (q.Quadrants[0] == nil || q.Quadrants[0].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
					q.subdivide()
				} else {
					q.Quadrants[0].Insert(point)
				}

			} else {
				log.Println("[ v ] \t\t The point is below of the center point!")
				if (q.Quadrants[2] == nil || q.Quadrants[2].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
					q.subdivide()
				} else {
					q.Quadrants[2].Insert(point)
				}
			}

		} else {
			log.Println("[ > ] \t\t The point is right of the center point!")

			// Test if the star is above or below the center point
			if point.C.Y > by && point.C.Y < by+bw {
				log.Println("[ ^ ] \t\t The point is above of the center point!")
				if (q.Quadrants[1] == nil || q.Quadrants[1].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
					q.subdivide()
				} else {
					q.Quadrants[1].Insert(point)
				}
			} else {
				log.Println("[ v ] \t\t The point is below of the center point!")
				if (q.Quadrants[3] == nil || q.Quadrants[3].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
					q.subdivide()
				} else {
					q.Quadrants[3].Insert(point)
				}
			}
		}
	}

	log.Printf("[>>>] Current Star [node]: %v", q.Star)
	for i := 0; i < 4; i++ {
		if q.Quadrants[i] != nil {
			log.Printf("[>>>] Current Star [%d]: %v", i, q.Quadrants[i].Star)
		} else {
			log.Printf("[>>>] Current Star [%d]: ", i)
		}
	}

	log.Printf("[   ] Tree after insertion: %v", q)
}

// subdivide subdivides the quadtree it is called on
func (q *Quadtree) subdivide() {
	// Toggle the leaf state: the node is not a leaf anymore
	q.Leaf = false

	// Get the "old" bounding box values
	oldCenterX := q.Boundary.Center.X
	oldCenterY := q.Boundary.Center.Y
	oldWidth := q.Boundary.Width

	// Calculate the bounding box values for the new quadrants
	newCenterNorthWest := Vec2{oldCenterX - (oldWidth / 4), oldCenterY + (oldWidth / 4)}
	newCenterNorthEast := Vec2{oldCenterX + (oldWidth / 4), oldCenterY + (oldWidth / 4)}
	newCenterSouthWest := Vec2{oldCenterX - (oldWidth / 4), oldCenterY - (oldWidth / 4)}
	newCenterSouthEast := Vec2{oldCenterX + (oldWidth / 4), oldCenterY - (oldWidth / 4)}

	// Calculate the new width
	newWidth := oldWidth / 2

	// Define the new bounding boxes using the values calculated above
	NorthWestBoundingBox := NewBoundingBox(newCenterNorthWest, newWidth)
	NorthEastBoundingBox := NewBoundingBox(newCenterNorthEast, newWidth)
	SouthWestBoundingBox := NewBoundingBox(newCenterSouthWest, newWidth)
	SouthEastBoundingBox := NewBoundingBox(newCenterSouthEast, newWidth)

	// Generate new quadrants using the new bounding boxes and assign them to the quadtree
	q.Quadrants[0] = NewQuadtree(NorthWestBoundingBox)
	q.Quadrants[1] = NewQuadtree(NorthEastBoundingBox)
	q.Quadrants[2] = NewQuadtree(SouthWestBoundingBox)
	q.Quadrants[3] = NewQuadtree(SouthEastBoundingBox)
}
