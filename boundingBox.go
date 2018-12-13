package structs

// BoundingBox is a struct defining the spatial outreach of a box
type BoundingBox struct {
	center Vec2    // Center of the box
	width  float64 // width of the box
}

func (b *BoundingBox) Width() float64 {
	return b.width
}

func (b *BoundingBox) SetWidth(width float64) {
	b.width = width
}

func (b *BoundingBox) Center() Vec2 {
	return b.center
}

func (b *BoundingBox) SetCenter(center Vec2) {
	b.center = center
}

func NewBoundingBox(center Vec2, halfDim float64) *BoundingBox {
	return &BoundingBox{center: center, width: halfDim}
}
