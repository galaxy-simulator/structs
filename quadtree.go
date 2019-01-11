package structs

type Node struct {
	Bounadry     BoundingBox // Spatial outreach of the quadtree
	CenterOfMass Vec2        // Center of mass of the cell
	TotalMass    float64     // Total mass of all the stars in the cell
	Depth        int         // Depth of the cell in the tree

	Star Star2D // The actual star

	// NW, NE, SW, SE
	Subtrees [4]*Node // The child subtrees
}
