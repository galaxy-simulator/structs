package structs

import (
	"fmt"
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
			starsize := 2

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

// DrawGalaxy draws the given quadtree to the given output path
func (q Quadtree) DrawGalaxy(outpath string) {
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

// GeneratePrintTree generates forest code for drawing a tree
func (q Quadtree) GeneratePrintTree(depth int) string {
	returnString := ""
	if q.Star != (Star2D{}) {
		returnString += "[a"
		fmt.Printf("[a")
	} else {
		returnString += "["
		fmt.Printf("[")
	}

	for i := 0; i < 4; i++ {
		if q.Quadrants[i] != nil {
			returnString += fmt.Sprintf("[%d]", depth)
			returnString += q.Quadrants[i].GeneratePrintTree(depth + 1)
		}
	}

	// ok, the reason the final image will only show the nodes in the leaf is, that in the latex
	// forest package that I use, trees must be drawn like this: [a[b]] and not like this: [[b]a].
	// [[b]a] == [[b]]. So there might be a lot of zeros, but that's ok!
	if q.Star != (Star2D{}) {
		returnString += "a]"
		fmt.Printf("a]")
	} else {
		returnString += "]"
		fmt.Printf("]")
	}

	return returnString
}

// DrawTree returns a valid LaTeX Document as a string drawing the quadtree it is called on using the forest package
func (q Quadtree) DrawTree() string {
	s1 := `\documentclass{article}
\usepackage{tikz}
\usepackage{forest}
\begin{document}
\begin{forest}
for tree={circle,draw, s sep+=0.25em}`

	s2 := q.GeneratePrintTree(0)

	s3 := `\end{forest}
\end{document}`

	return fmt.Sprintf("%s\n%s\n%s\n", s1, s2, s3)
}
