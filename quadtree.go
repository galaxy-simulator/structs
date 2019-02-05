// quadtree.go defines nodes and trees storing galaxies
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

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"sync"
)

// Node defines a node in the tree storing the galaxy
type Node struct {
	Boundary     BoundingBox // Spatial outreach of the quadtree
	CenterOfMass Vec2        // Center of mass of the cell
	TotalMass    float64     // Total mass of all the stars in the cell
	Depth        int         // Depth of the cell in the tree

	Star Star2D // The actual star

	// NW, NE, SW, SE
	Subtrees [4]*Node // The child subtrees
}

// NewRoot returns a pointer to a node defined as a root node. It taks the with of the BoundingBox as an argument
// resulting in a node that should (in theory) fit the whole galaxy if defined correctly.
func NewRoot(BoundingBoxWidth float64) *Node {
	return &Node{
		Boundary: BoundingBox{
			Center: Vec2{0, 0},
			Width:  BoundingBoxWidth,
		},
		CenterOfMass: Vec2{},
		TotalMass:    0,
		Depth:        0,
		Star:         Star2D{},
		Subtrees:     [4]*Node{},
	}
}

// NewNode creates a new new node using the given bounding box
func NewNode(bounadry BoundingBox) *Node {
	return &Node{Boundary: bounadry}
}

// Subdivide the tree
func (n *Node) Subdivide() {

	// define new values defining the new BoundaryBoxes
	newBoundaryWidth := n.Boundary.Width / 2
	newBoundaryPosX := n.Boundary.Center.X + (newBoundaryWidth / 2)
	newBoundaryPosY := n.Boundary.Center.Y + (newBoundaryWidth / 2)
	newBoundaryNegX := n.Boundary.Center.X - (newBoundaryWidth / 2)
	newBoundaryNegY := n.Boundary.Center.Y - (newBoundaryWidth / 2)

	// define the new Subtrees
	n.Subtrees[0] = NewNode(BoundingBox{Vec2{newBoundaryNegX, newBoundaryPosY}, newBoundaryWidth})
	n.Subtrees[1] = NewNode(BoundingBox{Vec2{newBoundaryPosX, newBoundaryPosY}, newBoundaryWidth})
	n.Subtrees[2] = NewNode(BoundingBox{Vec2{newBoundaryNegX, newBoundaryNegY}, newBoundaryWidth})
	n.Subtrees[3] = NewNode(BoundingBox{Vec2{newBoundaryPosX, newBoundaryNegY}, newBoundaryWidth})
}

// Insert inserts the given star into the Node or the tree it is called on
func (n *Node) Insert(star Star2D) error {
	var mutex = &sync.Mutex{}
	mutex.Lock()

	// if the subtree does not contain a node, insert the star
	if n.Star == (Star2D{}) {
		// if a subtree is present, insert the star into that subtree
		if n.Subtrees != [4]*Node{} {
			QuadrantBlocking := star.getRelativePositionInt(n.Boundary)
			err := n.Subtrees[QuadrantBlocking].Insert(star)
			if err != nil {
				fmt.Println(err)
			}

			// directly insert the star into the node
		} else {
			n.Star = star
			return nil
		}

		// Move the star blocking the slot into it's subtree using a recursive call on this function
		// and add the star to the slot
	} else {
		// if the node does not all ready have child nodes, subdivide it
		if n.Subtrees == ([4]*Node{}) {
			n.Subdivide()
		}

		// Insert the blocking star into it's subtree
		QuadrantBlocking := n.Star.getRelativePositionInt(n.Boundary)
		err := n.Subtrees[QuadrantBlocking].Insert(n.Star)
		if err != nil {
			fmt.Println(err)
		}
		n.Star = Star2D{}

		// Insert the blocking star into it's subtree
		QuadrantBlockingNew := star.getRelativePositionInt(n.Boundary)
		err = n.Subtrees[QuadrantBlockingNew].Insert(star)
		if err != nil {
			fmt.Println(err)
		}
		star = Star2D{}
	}

	mutex.Unlock()

	// fmt.Println("Done inserting %v, the tree looks like this: %v", star, n)
	return nil
}

// GenForestTree draws the subtree it is called on. If there is a star inside of the root node, the node is drawn
// The method returns a string depicting the tree in latex forest structure
func (n Node) GenForestTree(node *Node) string {

	returnstring := "["

	// if there is a star in the node, add the stars coordinates to the return string
	if n.Star != (Star2D{}) {
		returnstring += fmt.Sprintf("%.0f %.0f", n.Star.C.X, n.Star.C.Y)
	}

	// iterate over all the subtrees and call the GenForestTree method on the subtrees containing children
	for i := 0; i < len(n.Subtrees); i++ {
		if n.Subtrees[i] != nil {
			returnstring += n.Subtrees[i].GenForestTree(n.Subtrees[i])
		} else {
			returnstring += "[]"
		}
	}

	// Post-tree brace
	returnstring += "]"

	return returnstring
}

// DrawTreeLaTeX writes the tree it is called on to a texfile defined by the outpath parameter and
// calls lualatex to build the tex-file
func (n Node) DrawTreeLaTeX(outpath string) {
	// define all the stuff in front of the tree
	preamble := `\documentclass{article}
\usepackage{tikz}
\usepackage{forest}
\usepackage{adjustbox}

\begin{document}

\begin{adjustbox}{max size={\textwidth}{\textheight}}
\begin{forest}
for tree={,draw, s sep+=0.25em}
`

	// define all the stuff after the tree
	poststring := `
\end{forest}
\end{adjustbox}

\end{document}
`

	// combine all the strings
	data := []byte(fmt.Sprintf("%s%s%s", preamble, n.GenForestTree(&n), poststring))

	// write them to a file
	writeerr := ioutil.WriteFile(outpath, data, 0644)
	if writeerr != nil {
		panic(writeerr)
	}

	// build the pdf
	cmd := exec.Command("lualatex", outpath)
	runerr := cmd.Run()
	if runerr != nil {
		panic(runerr)
	}
}

// GetAllStars returns all the stars in the tree it is called on in an array
func (n Node) GetAllStars() []Star2D {

	// define a list to store the stars
	listOfNodes := []Star2D{}

	// if there is a star in the node, append the star to the list
	if n.Star != (Star2D{}) {
		listOfNodes = append(listOfNodes, n.Star)
	}

	// iterate over all the subtrees
	for i := 0; i < len(n.Subtrees); i++ {
		if n.Subtrees[i] != nil {

			// insert all the stars from the subtrees into the list of nodes
			for _, star := range n.Subtrees[i].GetAllStars() {
				listOfNodes = append(listOfNodes, star)
			}
		}
	}

	return listOfNodes
}

// CalcCenterOfMass calculates the center of mass for every node in the tree
func (n *Node) calcCenterOfMass() Vec2 {

	nominatorX := 0.0
	denominatorX := 0.0

	nominatorY := 0.0
	denominatorY := 0.0

	// if the subtrees are not empty
	if n.Subtrees != ([4]*Node{}) {
		for _, star := range n.Subtrees {
			fmt.Println(star)
		}
		for i := 0; i < len(n.Subtrees); i++ {
			nominatorX += n.Subtrees[i].calcCenterOfMass().X * n.Subtrees[i].TotalMass
			denominatorX += n.Subtrees[i].TotalMass

			nominatorY += n.Subtrees[i].calcCenterOfMass().Y * n.Subtrees[i].TotalMass
			denominatorY += n.Subtrees[i].TotalMass
		}
	}

	if n.Star != (Star2D{}) {
		n.CenterOfMass = n.Star.C
		fmt.Println(n.Star)
	}

	comX := nominatorX / denominatorX
	comY := nominatorY / denominatorY

	n.CenterOfMass = Vec2{comX, comY}

	return n.CenterOfMass
}

// CalcCenterOfMass calculates the center of mass for every node in the tree
func (n *Node) CalcCenterOfMass() Vec2 {
	tree := n.GenForestTree(n)
	fmt.Println(tree)
	centerOfMass := n.calcCenterOfMass()
	return centerOfMass
}

// CalcTotalMass calculates the total mass for every node in the tree
func (n *Node) CalcTotalMass() float64 {

	// if the subtrees are not empty
	if n.Subtrees != ([4]*Node{}) {
		for i := 0; i < len(n.Subtrees); i++ {
			n.TotalMass += n.Subtrees[i].CalcTotalMass()
		}
	}

	// if the star in the subtree is not empty
	if n.Star != (Star2D{}) {
		n.TotalMass += n.Star.M
	}

	return n.TotalMass
}

// CalcAllForces calculates the force acting in between the given star and all the other stars using the given theta.
// It gets all the other stars from the root node it is called on
func (n Node) CalcAllForces(star Star2D, theta float64) Vec2 {
	log.SetOutput(os.Stderr)

	// initialize a variable storing the overall force
	var localForce Vec2 = Vec2{}

	// calculate the local theta
	var tmpX float64 = math.Pow(star.C.X-n.Star.C.X, 2)
	var tmpY float64 = math.Pow(star.C.Y-n.Star.C.Y, 2)
	var distance float64 = math.Sqrt(tmpX + tmpY)

	var localtheta float64 = n.Boundary.Width / distance

	// if the subtree is not empty...
	if n.Subtrees != ([4]*Node{}) {

		// if the local theta is smaller than the given theta threshold...
		if localtheta < theta {
			// don't recurse further into the tree
			// calculate the forces in between the star and the node

			// define a new star using the center of mass of the new stars
			nodeStar := Star2D{
				C: Vec2{
					X: n.CenterOfMass.X,
					Y: n.CenterOfMass.Y,
				},
				V: Vec2{
					X: 0,
					Y: 0,
				},
				M: n.TotalMass,
			}

			// if the star is not equal to the node star, calculate the forces
			if star != nodeStar {

				// calculate the force on the individual star
				force := CalcForce(star, nodeStar)
				localForce.X += force.X
				localForce.Y += force.Y
			}

			// the local theta is bigger than the given theta -> recurse deeper
		} else {

			// iterate over all the subtrees
			for i := 0; i < len(n.Subtrees); i++ {
				force := n.Subtrees[i].CalcAllForces(star, theta)
				localForce.X += force.X
				localForce.Y += force.Y
			}
		}

		// if the subtree is empty
	} else {

		// make sure the star in the subtree is not empty
		if n.Star != (Star2D{}) {

			// if the star is not the star on which the forces should be calculated
			if star != n.Star {

				// calculate the forces acting on the star
				force := CalcForce(star, n.Star)
				localForce.X += force.X
				localForce.Y += force.Y
			}
		}
	}

	// return the overall acting force
	return localForce
}

// CalcForce calculates the force exerted on s1 by s2 and returns a vector representing that force
func CalcForce(s1 Star2D, s2 Star2D) Vec2 {
	G := 6.6726 * math.Pow(10, -11)

	// calculate the force acting
	var combinedMass float64 = s1.M * s2.M
	var distance float64 = math.Sqrt(math.Pow(math.Abs(s1.C.X-s2.C.X), 2) + math.Pow(math.Abs(s1.C.Y-s2.C.Y), 2))

	var scalar float64 = G * ((combinedMass) / math.Pow(distance, 2))

	// define a unit vector pointing from s1 to s2
	var vector Vec2 = Vec2{s2.C.X - s1.C.X, s2.C.Y - s1.C.Y}
	var UnitVector Vec2 = Vec2{vector.X / distance, vector.Y / distance}

	// multiply the vector with the force to get a vector representing the force acting
	var force Vec2 = UnitVector.Multiply(scalar)

	// return the force exerted on s1 by s2
	return force
}
