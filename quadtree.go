package structs

import (
	"log"
)

// Quadtree defines a quadtree and it's nodes recursively
type Quadtree struct {
	Boundary     BoundingBox `json:"boundary"`     // Spatial outreach of the quadtree
	CenterOfMass Vec2        `json:"CenterOfMass"` // Center of mass of the cell
	TotalMass    float64     `json:"totalMass"`    // Total mass of the cell
	Depth        int         `json:"depth"`        // Depth of the cell in the quadtree
	Star         Star2D      `json:"star"`         // Star inside the cell
	Leaf         bool        `json:"Leaf"`         // Quadtree is a leaf or not

	// NW, NE, SW, SE
	Quadrants [4]*Quadtree `json:"Quadrants"` // List of quadtrees representing individual Quadrants
}

type object struct {
	position Vec2
	mass     float64
	size     float64
	velocity Vec2
}

type node struct {
	father node
	Nodes  []*object
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

// Insert inserts the given point into the quadtree the method is called on
// func (q *Quadtree) Insert(point Star2D) {
// 	log.Printf("[   ] Inserting point %v into the tree %v", point, q)
//
// 	// prints the stars inside of the leaves of the current node
// 	log.Printf("[>>>] - Current Star [node]: %v", q.Star)
// 	for i := 0; i < 4; i++ {
// 		if q.Quadrants[i] != nil {
// 			log.Printf("[>>>] - Current Star [%d]: %v", i, q.Quadrants[i].Star)
// 		} else {
// 			log.Printf("[>>>] - Current Star [%d]: ", i)
// 		}
// 	}
//
// 	// create shortcuts for the various bounding box variables
// 	bx := q.Boundary.Center.X
// 	by := q.Boundary.Center.Y
// 	bw := q.Boundary.Width
// 	log.Printf("[ ~ ] \t Bounding Box X: %f", bx)
// 	log.Printf("[ ~ ] \t Bounding Box Y: %f", by)
// 	log.Printf("[ ~ ] \t Bounding Box Width: %f", bw)
//
// 	// Insert the given star into the tree
// 	// Case 1: There is no star inside of the node
// 	if q.Star == (Star2D{Vec2{}, Vec2{}, 0}) {
// 		log.Printf("[   ] There was no star inside of the node -> inserting directly")
// 		q.Star = point
//
// 		// if the star is not a leaf, try to insert the star into the correct leaf
// 		// if there all ready is a star inside of the leaf, subdivide that leaf and insert the two nodes recursively
// 		// into that leaf
// 		// TODO: implement the comment above
//
// 	// Case 2: There is all ready a star inside of the node
// 	} else {
// 		log.Printf("[ ! ] There is allready a star inside of the node -> subdividing")
//
// 		// Test if the star is left or right of the center point
// 		if point.C.X < bx && point.C.X > bx-bw {
// 			log.Println("[<  ] \t\t The point is left of the center point!")
//
// 			// Test if the star is above or below the center point
// 			if point.C.Y > by && point.C.Y < by+bw {
// 				log.Println("[ ^ ] \t\t The point is above of the center point!")
//
// 				// Subdivide if...
// 				// ... the quadrant does not contain a node yet or if there is all ready a node in the quadrant
// 				// ... the quadrant is a leaf
// 				if (q.Quadrants[0] == nil || q.Quadrants[0].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
// 					q.subdivide()
// 				} else {
// 					q.Quadrants[0].Insert(point)
// 				}
//
// 			} else {
// 				log.Println("[ v ] \t\t The point is below of the center point!")
// 				if (q.Quadrants[2] == nil || q.Quadrants[2].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
// 					q.subdivide()
// 				} else {
// 					q.Quadrants[2].Insert(point)
// 				}
// 			}
//
// 		} else {
// 			log.Println("[ > ] \t\t The point is right of the center point!")
//
// 			// Test if the star is above or below the center point
// 			if point.C.Y > by && point.C.Y < by+bw {
// 				log.Println("[ ^ ] \t\t The point is above of the center point!")
// 				if (q.Quadrants[1] == nil || q.Quadrants[1].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
// 					q.subdivide()
// 					q.Quadrants[1].Insert(point)
// 				} else {
// 					q.Quadrants[1].Insert(point)
// 				}
// 			} else {
// 				log.Println("[ v ] \t\t The point is below of the center point!")
// 				if (q.Quadrants[3] == nil || q.Quadrants[3].Star != (Star2D{Vec2{}, Vec2{}, 0})) && q.Leaf == true {
// 					q.subdivide()
// 				} else {
// 					q.Quadrants[3].Insert(point)
// 				}
// 			}
// 		}
// 	}
//
// 	log.Printf("[>>>] Current Star [node]: %v", q.Star)
// 	for i := 0; i < 4; i++ {
// 		if q.Quadrants[i] != nil {
// 			log.Printf("[>>>] Current Star [%d]: %v", i, q.Quadrants[i].Star)
// 		} else {
// 			log.Printf("[>>>] Current Star [%d]: ", i)
// 		}
// 	}
//
// 	// log.Printf("[   ] Tree after insertion: %v", *q)
// 	// q.print()
// }

// is Leaf returns true if the given node is a leaf (has no children) and false if it has children
func (q *Quadtree) isLeaf() bool {
	if q.Quadrants == ([4]*Quadtree{}) {
		return true
	} else {
		return false
	}
}

// noStar returns true if the current node does not contain a star and false if the node contains a star
func (q *Quadtree) noStar() bool {
	if q.Star == (Star2D{}) {
		return true
	} else {
		return false
	}
}

// NewInsert is a method inserting the given point (star) into the tree it is called on.
//
// The Method makes sure that all the points (stars) star are positioned in a leaf or get moved into a leaf
// so if there all ready is a star (B) in the node the star should be inserted into, (B) gets moved into the
// Leaf it belongs and the star that should be inserted gets inserted into it's place and is then shifted
// into the right child. This is repeated, until all the stars are inside of a leaf.
func (q *Quadtree) NewInsert(point Star2D) {
	log.Printf("[   ] Inserting the point %v into the tree", point)

	// so, using the example from the writeup, let's build this. we'll start with one of the more
	// simple cases: we've got an empty tree with an empty node and no children.
	// Inserting now works the following way:
	// The star gets inserted directly into the node, it now is a leaf node, because all of it's
	// children are non existent.
	//
	// Is a leaf (one of our goals \o/) and contains no star (our other goal \o/).
	if q.noStar() == true && q.Leaf == true {
		log.Println("[ + ] There is no other star in the node -> inserting directly")
		q.Star = point
		return
	}

	// Condition: cannot be inserted and into node that is a leaf because of another blocking star
	// Solution: insert the star blocking the node into the nodes subtree and then insert the star
	// into the nodes subtree as well, if the subtree is occupied by an other star, reinsert
	if q.noStar() == false && q.Leaf == true {
		q.insertHasStarIsLeaf(point)
		return
	}

	if q.noStar() == true && q.Leaf == false {
		q.insertHasNoStarIsNotLeaf(point)
		return
	}

	if q.noStar() == false && q.Leaf == false {
		q.insertHasNoStarIsNotLeaf(point)
		return
	}

	// // Let's continue with the case above: we've now got a tree with a star (A) in the space we want
	// // to insert the new star (B), so we need to shift (A) down into the tree and then insert (B)
	// // into the corresponding space so long, until they both are in leaves.
	// if q.noStar() == false {
	//
	// 	// We define A as the star that is all ready inside of the tree and B as the star that should
	// 	// be inserted
	// 	A := q.Star
	// 	B := point
	//
	// 	// The if conditions below evaluate in which of the cells A is.
	// 	// If there is no tree in that quadrant, the tree is subdivided and a 4 new quadrants are generated
	// 	// In the end, the star is inserted into that quadrant
	// 	if A.C.Y > by && A.C.Y < by+bw {
	// 		if A.C.Y > bx && A.C.Y < bx+bw {
	// 			if q.Quadrants[0] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[0].NewInsert(A)
	// 		} else {
	// 			if q.Quadrants[2] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[2].NewInsert(A)
	// 		}
	//
	// 	} else {
	// 		// Test if b is left or right of the x range
	// 		if A.C.Y > bx && A.C.Y < bx+bw {
	// 			if q.Quadrants[1] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[1].NewInsert(A)
	// 		} else {
	// 			if q.Quadrants[3] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[3].NewInsert(A)
	// 		}
	// 	}
	//
	// 	// So after inserting A into the tree, we still need to insert B:
	// 	if B.C.Y > by && B.C.Y < by+bw {
	// 		if B.C.Y > bx && B.C.Y < bx+bw {
	// 			if q.Quadrants[0] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[0].NewInsert(B)
	// 		} else {
	// 			if q.Quadrants[2] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[2].NewInsert(B)
	// 		}
	//
	// 	} else {
	// 		if B.C.Y > bx && B.C.Y < bx+bw {
	// 			if q.Quadrants[1] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[1].NewInsert(B)
	// 		} else {
	// 			if q.Quadrants[3] == nil {
	// 				q.subdivide()
	// 			}
	// 			q.Quadrants[3].NewInsert(B)
	// 		}
	// 	}
	// }

}

// insertHasStarIsLeaf inserts a given point into the quadtree it is called on.
// The condition is, that the node in which the star should be inserted into already contains a star, so the
// star that is all ready inside the node must be moved further into the tree before inserting the new star
func (q *Quadtree) insertHasStarIsLeaf(point Star2D) {
	q.GeneratePrintTree(0)

	bx := q.Boundary.Center.X
	by := q.Boundary.Center.Y
	bw := q.Boundary.Width

	log.Println("[ i ] Star blocking the node, but the node is a leaf")

	log.Println("[   ] Subdividing the node")
	// subdivide the node
	q.subdivide()

	log.Println("[   ] Inserting the star blocking the node into the nodes subtree")

	// insert the node blocking the tree into the newly created subtree
	A := q.Star
	if A.C.Y > by && A.C.Y < by+bw {
		log.Println("[<  ] \t\t The point is left of the center point!")
		if A.C.Y > bx && A.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			if q.Quadrants[0].Star != (Star2D{}) {
				q.Quadrants[0].insertHasStarIsLeaf(point)
			} else {
				q.Quadrants[0].NewInsert(A)
			}
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			if q.Quadrants[2].Star != (Star2D{}) {
				q.Quadrants[2].insertHasStarIsLeaf(point)
			} else {
				q.Quadrants[2].NewInsert(A)
			}
		}

	} else {
		log.Println("[ > ] \t\t The point is right of the center point!")
		// Test if b is left or right of the x range
		if A.C.Y > bx && A.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			if q.Quadrants[1].Star != (Star2D{}) {
				q.Quadrants[1].insertHasStarIsLeaf(point)
			} else {
				q.Quadrants[1].NewInsert(A)
			}
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			if q.Quadrants[3].Star != (Star2D{}) {
				q.Quadrants[3].insertHasStarIsLeaf(point)
			} else {
				q.Quadrants[3].NewInsert(A)
			}
		}
	}

	log.Println("[   ] Inserting the new star into the new subtree")

	// Insert the point into the newly created subtree
	B := point
	if B.C.Y > by && B.C.Y < by+bw {
		log.Println("[<  ] \t\t The point is left of the center point!")
		if B.C.Y > bx && B.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			q.Quadrants[0].NewInsert(B)
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			q.Quadrants[2].NewInsert(B)
		}

	} else {
		log.Println("[ > ] \t\t The point is right of the center point!")
		if B.C.Y > bx && B.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			q.Quadrants[1].NewInsert(B)
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			q.Quadrants[3].NewInsert(B)
		}
	}

	// retry to insert the point
	q.NewInsert(point)

	q.Star = Star2D{}
}

// insertHasNoStarIsNotLeaf inserts a given point into the quadtree it is called on.
// The condition is, that the node in which the star should be inserted into does not contain a star, but
// also isn't a leaf.
func (q *Quadtree) insertHasNoStarIsNotLeaf(point Star2D) {
	log.Println("[ i ] The node does not contain a star and is not a leaf")

	bx := q.Boundary.Center.X
	by := q.Boundary.Center.Y
	bw := q.Boundary.Width

	q.GeneratePrintTree(0)

	// Inserting the star into the subtree
	// If there is all ready a star in the slot, insert that star into it's subtree and then insert the star into
	// that subtree recursively
	A := point
	if A.C.Y > by && A.C.Y < by+bw {
		log.Println("[<  ] \t\t The point is left of the center point!")
		if A.C.Y > bx && A.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			if q.Quadrants[0].Star != (Star2D{}) {
				q.Quadrants[0].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[0].NewInsert(A)
			}
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			if q.Quadrants[2].Star != (Star2D{}) {
				q.Quadrants[2].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[2].NewInsert(A)
			}
		}

	} else {
		log.Println("[ > ] \t\t The point is right of the center point!")
		// Test if b is left or right of the x range
		if A.C.Y > bx && A.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			if q.Quadrants[1].Star != (Star2D{}) {
				q.Quadrants[1].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[1].NewInsert(A)
			}
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			if q.Quadrants[3].Star != (Star2D{}) {
				q.Quadrants[3].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[3].NewInsert(A)
			}
		}
	}

}

// insertHasStarIsNotLeaf inserts a given point into the quadtree it is called on.
// The condition is, that the node in which the star should be inserted into does contain a star, but
// also isn't a leaf.
func (q *Quadtree) insertHasStarIsNotLeaf(point Star2D) {
	log.Println("[ i ] The node contains a star and is not a leaf")

	bx := q.Boundary.Center.X
	by := q.Boundary.Center.Y
	bw := q.Boundary.Width

	q.GeneratePrintTree(0)

	// Inserting the star into the subtree
	// Because of there all ready being a star in the tree, the star in the tree get's inserted into the
	// subtree first, the point to be inserted is then also inserted into the subtree
	A := q.Star
	if A.C.Y > by && A.C.Y < by+bw {
		log.Println("[<  ] \t\t The point is left of the center point!")
		if A.C.Y > bx && A.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			if q.Quadrants[0].Star != (Star2D{}) {
				q.Quadrants[0].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[0].NewInsert(A)
			}
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			if q.Quadrants[2].Star != (Star2D{}) {
				q.Quadrants[2].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[2].NewInsert(A)
			}
		}

	} else {
		log.Println("[ > ] \t\t The point is right of the center point!")
		// Test if b is left or right of the x range
		if A.C.Y > bx && A.C.Y < bx+bw {
			log.Println("[ ^ ] \t\t The point is above of the center point!")
			if q.Quadrants[1].Star != (Star2D{}) {
				q.Quadrants[1].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[1].NewInsert(A)
			}
		} else {
			log.Println("[ v ] \t\t The point is below of the center point!")
			if q.Quadrants[3].Star != (Star2D{}) {
				q.Quadrants[3].NewInsert(q.Star)
				q.Quadrants[0].NewInsert(point)
			} else {
				q.Quadrants[3].NewInsert(A)
			}
		}
	}
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
