package structs

// BoundingBox is a struct defining the spatial outreach of a box
type BoundingBox struct {
	Center Vec2    // Center of the box
	Width  float64 // Width of the box
}

// NewBoundingBox returns a new Bounding Box using the centerpoint and the width given by the function parameters
func NewBoundingBox(center Vec2, width float64) BoundingBox {
	return BoundingBox{Center: center, Width: width}
}
