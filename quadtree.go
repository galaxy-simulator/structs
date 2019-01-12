package structs

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

type Node struct {
	Boundry      BoundingBox // Spatial outreach of the quadtree
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
		Boundry: BoundingBox{
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

// Create a new new node using the given bounding box
func NewNode(bounadry BoundingBox) *Node {
	return &Node{Boundry: bounadry}
}

// Subdivide the tree
func (n *Node) subdivide() {

	// define new values defining the new BoundaryBoxes
	newBoundaryWidth := n.Boundry.Width / 2
	newBoundaryPosX := n.Boundry.Center.X + (newBoundaryWidth / 2)
	newBoundaryPosY := n.Boundry.Center.Y + (newBoundaryWidth / 2)
	newBoundaryNegX := n.Boundry.Center.X - (newBoundaryWidth / 2)
	newBoundaryNegY := n.Boundry.Center.Y - (newBoundaryWidth / 2)

	// define the new Subtrees
	n.Subtrees[0] = NewNode(BoundingBox{Vec2{newBoundaryNegX, newBoundaryPosY}, newBoundaryWidth})
	n.Subtrees[1] = NewNode(BoundingBox{Vec2{newBoundaryPosX, newBoundaryPosY}, newBoundaryWidth})
	n.Subtrees[2] = NewNode(BoundingBox{Vec2{newBoundaryNegX, newBoundaryNegY}, newBoundaryWidth})
	n.Subtrees[3] = NewNode(BoundingBox{Vec2{newBoundaryPosX, newBoundaryNegY}, newBoundaryWidth})
}

// Insert inserts the given star into the Node or the tree it is called on
func (n *Node) Insert(star Star2D) {

	// prevent the function to recurse to deep into the tree
	if n.Boundry.Width < 0.1 {
		return
	}

	// if the subtree does not contain a node, insert the star
	if n.Star == (Star2D{}) {

		// if a subtree is present, insert the star into that subtree
		if n.Subtrees != [4]*Node{} {
			QuadrantBlocking := star.getRelativePositionInt(n.Boundry)
			n.Subtrees[QuadrantBlocking].Insert(star)

			// directly insert the star into the node
		} else {
			n.Star = star
		}

		// Move the star blocking the slot into it's subtree using a recursive call on this function
		// and add the star to the slot
	} else {

		// if the node does not all ready have child nodes, subdivide it
		if n.Subtrees == ([4]*Node{}) {
			n.subdivide()
		}

		// Insert the blocking star into it's subtree
		QuadrantBlocking := n.Star.getRelativePositionInt(n.Boundry)
		n.Subtrees[QuadrantBlocking].Insert(n.Star)
		n.Star = Star2D{}

		// Insert the blocking star into it's subtree
		QuadrantBlockingNew := star.getRelativePositionInt(n.Boundry)
		n.Subtrees[QuadrantBlockingNew].Insert(star)
		star = Star2D{}

	}
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
