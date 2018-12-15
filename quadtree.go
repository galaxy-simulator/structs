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
	//newquadtree := &Quadtree{
	//	Boundary: boundary,
	//}

	newquadtree := &Quadtree{
		Boundary: BoundingBox{
			Center: Vec2{
				X: boundary.Center.X,
				Y: boundary.Center.Y,
			},
			Width: boundary.Width,
		},
		CenterOfMass: Vec2{
			X: 0,
			Y: 0,
		},
		TotalMass: 0,
		Depth:     0,
		Star: Star2D{
			C: Vec2{
				X: 0,
				Y: 0,
			},
			V: Vec2{
				X: 0,
				Y: 0,
			},
			M: 0,
		},
		Leaf:      false,
		Quadrants: [4]*Quadtree{},
	}
	return newquadtree
}

// Insert inserts the given point into the quadtree the method is called on
func (q *Quadtree) Insert(point Star2D) {
	log.Printf("[ ] Inserting point %v into the tree %v", point, q)
	bx := q.Boundary.Center.X
	by := q.Boundary.Center.Y
	bw := q.Boundary.Width
	log.Printf("[~] \t Bounding Box X: %f", bx)
	log.Printf("[~] \t Bounding Box Y: %f", by)
	log.Printf("[~] \t Bounding Box Width: %f", bw)

	var empty bool = true // is the tree empty?
	for _, element := range q.Quadrants {

		// if one element is not empty
		if element != nil {
			empty = false
		}
	}

	empty = true

	if empty == true {
		log.Println("[ ] Subdividing the current tree")
		q.subdivide()
		log.Println("[+] Done Subdividing!")

		log.Printf("[~] \t point: %v\n", point)
		log.Printf("[~] \t quadrant: %v\n", q.Quadrants[0])

		if point.C.X < bx && point.C.X > bx-bw {
			// Left
			log.Println("[~] \t\t The point is left of the y-axis!")
			if point.C.Y > by && point.C.Y < by+bw {
				// Top Left
				log.Println("[~] \t\t The point is above of the x-axis!")
				log.Println("[ ] \t Inserting the point into the top left (NW) quadtree")
				q.Quadrants[0].Star = point
				log.Println("[+] \t DONE!")
			} else {
				// Bottom Left
				log.Println("[~] \t\t The point is below of the x-axis!")
				log.Println("[ ] \t Inserting the point into the bottom left (SW) quadtree")
				q.Quadrants[2].Star = point
				log.Println("[+] \t DONE!")
			}
		} else {
			// Right
			log.Println("[~] \t\t The point is right of the y-axis!")
			if point.C.Y > by && point.C.Y < by+bw {
				// Top Right
				log.Println("[~] \t\t The point is above of the x-axis!")
				log.Println("[ ] \t Inserting the point into the top right (NE) quadtree")
				q.Quadrants[1].Star = point
				log.Println("[+] \t DONE!")
			} else {
				// Bottom Right
				log.Println("[~] \t\t The point is below of the x-axis!")
				log.Println("[ ] \t Inserting the point into the bottom right (SE) quadtree")
				q.Quadrants[3].Star = point
				log.Println("[+] \t DONE!")
			}
		}
	}
}

// subdivide subdivides the quadtree it is called on
func (q *Quadtree) subdivide() {
	log.Println("[ ] Getting the current boundary")
	oldCenterX := q.Boundary.Center.X
	oldCenterY := q.Boundary.Center.Y
	oldWidth := q.Boundary.Width
	log.Printf("[~] \t oldCenterX: %f\n", oldCenterX)
	log.Printf("[~] \t oldCenterY: %f\n", oldCenterY)
	log.Printf("[~] \t oldWidth: %f\n", oldWidth)
	log.Println("[+] Done getting the current boundary!")

	log.Println("[ ] Defining the new centerpoints")
	newCenterNW := Vec2{oldCenterX - (oldWidth / 2), oldCenterY + (oldWidth / 2)}
	newCenterNE := Vec2{oldCenterX + (oldWidth / 2), oldCenterY + (oldWidth / 2)}
	newCenterSW := Vec2{oldCenterX - (oldWidth / 2), oldCenterY - (oldWidth / 2)}
	newCenterSE := Vec2{oldCenterX + (oldWidth / 2), oldCenterY - (oldWidth / 2)}
	log.Printf("[~] \t newCenterNW: %v\n", newCenterNW)
	log.Printf("[~] \t newCenterNE: %v\n", newCenterNE)
	log.Printf("[~] \t newCenterSW: %v\n", newCenterSW)
	log.Printf("[~] \t newCenterSE: %v\n", newCenterSE)
	log.Println("[+] Done defining the new centerpoints!")

	log.Println("[ ] Calculating th new width")
	log.Printf("[~] \t Old width: %f", oldWidth)
	newWidth := oldWidth / 2
	log.Printf("[~] \t New width: %f", newWidth)
	log.Println("[+] Done calculating the new width!")

	log.Println("[ ] Generating the new bounding boxes")
	NWboundingBox := NewBoundingBox(newCenterNW, newWidth)
	NEboundingBox := NewBoundingBox(newCenterNE, newWidth)
	SWboundingBox := NewBoundingBox(newCenterSW, newWidth)
	SEboundingBox := NewBoundingBox(newCenterSE, newWidth)
	log.Printf("[~] \t NW: %v", NWboundingBox)
	log.Printf("[~] \t NE: %v", NEboundingBox)
	log.Printf("[~] \t SW: %v", SWboundingBox)
	log.Printf("[~] \t SE: %v", SEboundingBox)
	log.Println("[+] Done generating the new bounding boxes!")

	log.Println("[ ] assigning the bounding boxes to the individual quadrants")
	log.Printf("[~] \t root quadtree: %v\n", q)
	log.Printf("[~] \t NW quadtree: %v\n", NewQuadtree(NWboundingBox))
	log.Printf("[~] \t NE quadtree: %v\n", NewQuadtree(NEboundingBox))
	log.Printf("[~] \t SW quadtree: %v\n", NewQuadtree(SWboundingBox))
	log.Printf("[~] \t SE quadtree: %v\n", NewQuadtree(SEboundingBox))
	q.Quadrants[0] = NewQuadtree(NWboundingBox)
	q.Quadrants[1] = NewQuadtree(NEboundingBox)
	q.Quadrants[2] = NewQuadtree(SWboundingBox)
	q.Quadrants[3] = NewQuadtree(SEboundingBox)
	log.Println("[+] Done assigning the bounding boxes to the individual quadrants!")
}
