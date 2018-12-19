package structs

import (
	"log"

	"github.com/fogleman/gg"
)

func initializePlot(imageWidth int, imageHeight int) *gg.Context {

	// define an new context using the given image width and height
	context := gg.NewContext(imageWidth, imageHeight)

	// set the background color to black
	context.SetRGB(0, 0, 0)
	context.Clear()
	context.SetRGB(1, 1, 1)

	// translate the coordinate origin to the midpoint of the image
	context.Translate(float64(imageWidth/2), float64(imageHeight/2))

	return context
}

// drawQuadtree draws a given quadtree and the stars in it recursively to the given canvas
func drawQuadtree(context *gg.Context, q Quadtree) {

	// find out if the node is a leaf find out if the node is a leaf find out if the node is a leaf find out if the node is a leaf
	var draw bool = false
	for i := 0; i < 4; i++ {
		if q.Quadrants[i] == nil {
			draw = true
		}
	}

	// draw the bounding box and the star if the node is a leaf
	if draw == true {

		// Don't draw nonexistent stars
		if q.Star != (Star2D{}) {

			// define the current star
			x := q.Star.C.X
			y := q.Star.C.Y
			starsize := 1

			// set the color of the stars to green
			context.SetRGB(0, 1, 1)
			context.DrawPoint(x, y, float64(starsize))
			context.Fill()
			// context.DrawString(fmt.Sprintf("(%f, %f)", x, y), x, y)
			context.Stroke()
			log.Printf("[***] Drawing Star (%f, %f)", x, y)
		}

		// define the bounding box
		boundingx := q.Boundary.Center.X
		boundingy := q.Boundary.Center.Y
		boundingw := q.Boundary.Width

		// bottom left corner
		contextx := boundingx - (boundingw / 2)
		contexty := boundingy - (boundingw / 2)

		// draw the rectangle
		context.SetRGB(1, 1, 1)
		context.DrawRectangle(contextx, contexty, boundingw, boundingw)
		context.Stroke()

		log.Printf("[***] Drawing Box ((%f, %f), %f)", contextx, contexty, boundingw)
	}

	// draw all the other trees recursively...
	for i := 0; i < 4; i++ {
		// ... but only if they exist
		if q.Quadrants[i] != nil {
			drawQuadtree(context, *q.Quadrants[i])
		}
	}
}

// saveImage saves the given context to the given path as a png
func saveImage(context *gg.Context, outpath string) {
	savePngError := context.SavePNG(outpath)
	if savePngError != nil {
		panic(savePngError)
	}
}

// Draw draws the given quadtree to the given output path
func (q Quadtree) Draw(outpath string) {
	log.Printf("Drawing the quadtree to %s", outpath)

	// define the image dimensions
	imageWidth := 1024 * 8
	imageHeight := 1024 * 8

	// define a new context to draw on
	context := initializePlot(imageWidth, imageHeight)

	// first recursive call of drawQuadtree
	drawQuadtree(context, q)

	// save the context to the given output path
	saveImage(context, outpath)
}
