package structs

import (
	"fmt"
	"reflect"
	"testing"
)

// The example below creates a new root node at (0, 0) with the given width.
// The width given in this case is 100.
func ExampleNewRoot() {
	root := NewRoot(100)
	fmt.Printf("%v\n", root)
	// Output: &{{{0 0} 100} {0 0} 0 0 {{0 0} {0 0} 0} [<nil> <nil> <nil> <nil>]}
}

func TestNewRoot(t *testing.T) {
	type args struct {
		BoundingBoxWidth float64
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "New root at (0, 0) with a width of 100",
			args: args{
				BoundingBoxWidth: 100,
			},
			want: &Node{
				Boundary: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 100,
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
				Subtrees: [4]*Node{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoot(tt.args.BoundingBoxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

// The example below creates a new node using the given bounding box
func ExampleNewNode() {
	newNode := NewNode(BoundingBox{
		Center: Vec2{
			X: 25,
			Y: 25,
		},
		Width: 50,
	})
	fmt.Printf("%v\n", newNode)
	// Output: &{{{25 25} 50} {0 0} 0 0 {{0 0} {0 0} 0} [<nil> <nil> <nil> <nil>]}
}

func TestNewNode(t *testing.T) {
	type args struct {
		bounadry BoundingBox
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Return a new node",
			args: args{
				bounadry: BoundingBox{
					Center: Vec2{
						X: 3,
						Y: 15,
					},
					Width: 100,
				},
			},
			want: &Node{
				Boundary: BoundingBox{
					Center: Vec2{
						X: 3,
						Y: 15,
					},
					Width: 100,
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
				Subtrees: [4]*Node{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.bounadry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// The example below subdivides the node it is called on.
// The function inserts four pointers pointing to other nodes into the subtree array representing the quadrants.
//
// The Code below prints the four subtrees that where generated:
func ExampleNode_Subdivide() {
	root := NewRoot(100)
	root.Subdivide()
	for i := 0; i < 4; i++ {
		fmt.Printf("%v\n", root.Subtrees[i])
	}
	// Output:
	// &{{{-25 25} 50} {0 0} 0 0 {{0 0} {0 0} 0} [<nil> <nil> <nil> <nil>]}
	// &{{{25 25} 50} {0 0} 0 0 {{0 0} {0 0} 0} [<nil> <nil> <nil> <nil>]}
	// &{{{-25 -25} 50} {0 0} 0 0 {{0 0} {0 0} 0} [<nil> <nil> <nil> <nil>]}
	// &{{{25 -25} 50} {0 0} 0 0 {{0 0} {0 0} 0} [<nil> <nil> <nil> <nil>]}
}

func TestNode_Subdivide(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Subdivide a tree from [100] to [[50][50][50][50]]",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 100,
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
				Subtrees: [4]*Node{
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: 25,
							},
							Width: 50,
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
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: 25,
							},
							Width: 50,
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
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: -25,
							},
							Width: 50,
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
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: -25,
							},
							Width: 50,
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
						Subtrees: [4]*Node{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			n.Subdivide()
		})
	}
}

// Insert a star into a tree.
// If the star cannot be inserted (e.g. recursion depth too deep), Insert() returns an error
func ExampleNode_Insert() {

	// Initialize a tree and a star
	root := NewRoot(100)
	star := Star2D{
		C: Vec2{
			X: 12,
			Y: 34,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 0,
	}

	// insert the star into the tree
	err := root.Insert(star)

	// handle potential errors
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", root)

	// Output:
	// Direct insert of (12.000000, 34.000000)
	// &{{{0 0} 100} {0 0} 0 0 {{12 34} {0 0} 0} [<nil> <nil> <nil> <nil>]}
}

// Insert two stars that are very close to each other into the tree.
// A problem arises: the tree has to be subdivided so often, that the insert function
// raises a recursion depth error
func ExampleNode_Insert_error() {

	// Initialize a tree and two
	root := NewRoot(100)
	star1 := Star2D{
		C: Vec2{
			X: 5,
			Y: 5,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 0,
	}
	star2 := Star2D{
		C: Vec2{
			X: 5.000000000000001,
			Y: 5.000000000000001,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 0,
	}

	// insert the first star into the tree
	err := root.Insert(star1)
	// handle potential errors
	if err != nil {
		panic(err)
	}

	// insert the second star into the tree
	err = root.Insert(star2)
	// handle potential errors
	if err != nil {
		panic(err)
	}

	// Output:
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Direct insert of (5.000000, 5.000000)
	// Could not insert star (5.000000, 5.000000) (recursion limit reached)
	//
	// Could not insert star (5.000000, 5.000000) (recursion limit reached)
}

func TestNode_Insert(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	type args struct {
		star Star2D
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Inserting a single star into a previously empty galaxy",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 100,
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
				Subtrees: [4]*Node{},
			},
			args: args{
				star: Star2D{
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
			},
			wantErr: false,
		},
		{
			name: "Inserting a single star into a galaxy all ready containing a star",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 100,
				},
				CenterOfMass: Vec2{
					X: 0,
					Y: 0,
				},
				TotalMass: 0,
				Depth:     0,
				Star: Star2D{
					C: Vec2{
						X: 2,
						Y: 3,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
				Subtrees: [4]*Node{},
			},
			args: args{
				star: Star2D{
					C: Vec2{
						X: 10,
						Y: 20,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Inserting a single star onto the boundary limit",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 100,
				},
				CenterOfMass: Vec2{
					X: 0,
					Y: 0,
				},
				TotalMass: 0,
				Depth:     0,
				Star: Star2D{
					C: Vec2{
						X: 10,
						Y: 20,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
				Subtrees: [4]*Node{},
			},
			args: args{
				star: Star2D{
					C: Vec2{
						X: 25,
						Y: 25,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 20,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			if err := n.Insert(tt.args.star); (err != nil) != tt.wantErr {
				t.Errorf("Node.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Generate a tree using the LaTeX forest tree notation
// This is a minimal example using only a root node
func ExampleNode_GenForestTree() { // Create a new root
	root := NewRoot(100)
	root.Star = Star2D{
		C: Vec2{
			X: 10,
			Y: 20,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 20,
	}

	// generate the tree
	forestTree := root.GenForestTree(root)
	fmt.Println(forestTree)
	// Output:
	// [10 20[][][][]]
}

// Generate a tree using the LaTeX forest tree notation.
// This is an example displaying a bigger tree.
func ExampleNode_GenForestTree_deepTree() {
	// Create a new root
	root := NewRoot(100)

	// Subdivide the root multiple times
	root.Subdivide()
	root.Subtrees[1].Subdivide()

	// Insert a star into the tree
	root.Subtrees[1].Star = Star2D{
		C: Vec2{
			X: 20,
			Y: 30,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 20,
	}

	// Generate the tree
	forestTree := root.GenForestTree(root)
	fmt.Println(forestTree)
	// Output:
	// [[[][][][]][20 30[[][][][]][[][][][]][[][][][]][[][][][]]][[][][][]][[][][][]]]
}

func TestNode_GenForestTree(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Create a forest from a single node with four subnodes",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 100,
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
				Subtrees: [4]*Node{},
			},
			args: args{
				node: &Node{
					Boundary: BoundingBox{
						Center: Vec2{
							X: 0,
							Y: 0,
						},
						Width: 100,
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
					Subtrees: [4]*Node{},
				},
			},
			want: "[[][][][]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			if got := n.GenForestTree(tt.args.node); got != tt.want {
				t.Errorf("Node.GenForestTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Draws the tree to a pdf using lualatex for building the pdf.
// (Luatex is used, because pdflatex apparently cannot handle such deep recursion depths)
func ExampleNode_DrawTreeLaTeX() {
	// create a new root node
	root := NewRoot(100)

	// write the LaTeX to out.tex and build the tex using luatex
	root.DrawTreeLaTeX("out.tex")
	// Output:
}

func TestNode_DrawTreeLaTeX(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	type args struct {
		outpath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Draw a tree using latex",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 0,
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
				Subtrees: [4]*Node{},
			},
			args: args{
				outpath: "tree.tex",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			n.DrawTreeLaTeX(tt.args.outpath)
		})
	}
}

// GetAllStars gets all the stars from a selected tree.
// In the example below, an empty node is generated, subdivided and two stars are inserted into it.
// Then, a list of all stars (the two that were previously inserted) gets generated and printed.
func ExampleNode_GetAllStars() {
	// Define a new root node
	root := NewRoot(100)

	// Subdivide the root
	root.Subdivide()

	// Insert two stars into the tree
	root.Subtrees[1].Star = Star2D{
		C: Vec2{
			X: 10,
			Y: 20,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 0,
	}
	root.Subtrees[3].Star = Star2D{
		C: Vec2{
			X: 30,
			Y: 40,
		},
		V: Vec2{
			X: 0,
			Y: 0,
		},
		M: 0,
	}

	// Get the stars from the tree
	starsList := root.GetAllStars()

	// Print all the stars in the list
	for _, star := range starsList {
		fmt.Println(star)
	}
	// Output:
	// {{10 20} {0 0} 0}
	// {{30 40} {0 0} 0}
}

func TestNode_GetAllStars(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []Star2D
	}{
		{
			name: "",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 0,
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
				Subtrees: [4]*Node{
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: 25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 10,
								Y: 20,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: 25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 30,
								Y: 40,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: -25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 50,
								Y: 60,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: -25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 70,
								Y: 80,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
				},
			},
			want: []Star2D{
				{
					C: Vec2{
						X: 10,
						Y: 20,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
				{
					C: Vec2{
						X: 30,
						Y: 40,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
				{
					C: Vec2{
						X: 50,
						Y: 60,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
				{
					C: Vec2{
						X: 70,
						Y: 80,
					},
					V: Vec2{
						X: 0,
						Y: 0,
					},
					M: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			if got := n.GetAllStars(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.GetAllStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

// CalcCenterOfMass calculates the center of mass of the node it is called on.
// In the example below, the Tree contains two stars with equal mass (10): (20, 30) and (-20, -30).
func ExampleNode_CalcCenterOfMass() {
	root := NewRoot(100)

	root.Subdivide()

	star1 := NewStar2D(Vec2{0, 0}, Vec2{0, 0}, 10)
	star2 := NewStar2D(Vec2{3, 3}, Vec2{0, 0}, 10)

	// Insert the stars into the tree
	// There will be no error handling here, we'll assume that everything goes right..
	_ = root.Insert(star1)
	_ = root.Insert(star2)

	root.CalcTotalMass()

	centerOfMass := root.CalcCenterOfMass()
	fmt.Println(centerOfMass)
	// Output:
}

func TestNode_CalcCenterOfMass(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec2
	}{
		{
			name: "Center of mass inbetween ",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 0,
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
				Subtrees: [4]*Node{
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: 25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 10,
								Y: 20,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: 25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 30,
								Y: 40,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: -25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 50,
								Y: 60,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: -25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 70,
								Y: 80,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 0,
						},
						Subtrees: [4]*Node{},
					},
				},
			},
			want: Vec2{
				X: 0,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			if got := n.CalcCenterOfMass(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.CalcCenterOfMass() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
func ExampleNode_CalcTotalMass() {

}

func TestNode_CalcTotalMass(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Calculating the total mass two stars",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 0,
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
				Subtrees: [4]*Node{
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: 25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 10,
								Y: 20,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 42,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: 25,
							},
							Width: 50,
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
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: -25,
								Y: -25,
							},
							Width: 50,
						},
						CenterOfMass: Vec2{
							X: 0,
							Y: 0,
						},
						TotalMass: 0,
						Depth:     0,
						Star: Star2D{
							C: Vec2{
								X: 30,
								Y: 40,
							},
							V: Vec2{
								X: 0,
								Y: 0,
							},
							M: 24,
						},
						Subtrees: [4]*Node{},
					},
					{
						Boundary: BoundingBox{
							Center: Vec2{
								X: 25,
								Y: -25,
							},
							Width: 50,
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
						Subtrees: [4]*Node{},
					},
				},
			},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			if got := n.CalcTotalMass(); got != tt.want {
				t.Errorf("Node.CalcTotalMass() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
func ExampleNode_CalcAllForces() {

}

func TestNode_CalcAllForces(t *testing.T) {
	type fields struct {
		Boundry      BoundingBox
		CenterOfMass Vec2
		TotalMass    float64
		Depth        int
		Star         Star2D
		Subtrees     [4]*Node
	}
	type args struct {
		star  Star2D
		theta float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{
			name: "single star",
			fields: fields{
				Boundry: BoundingBox{
					Center: Vec2{
						X: 0,
						Y: 0,
					},
					Width: 0,
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
				Subtrees: [4]*Node{},
			},
			args: args{
				star: Star2D{
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
				theta: 0,
			},
			want: Vec2{
				X: 0,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Boundary:     tt.fields.Boundry,
				CenterOfMass: tt.fields.CenterOfMass,
				TotalMass:    tt.fields.TotalMass,
				Depth:        tt.fields.Depth,
				Star:         tt.fields.Star,
				Subtrees:     tt.fields.Subtrees,
			}
			if got := n.CalcAllForces(tt.args.star, tt.args.theta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.CalcAllForces() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
func ExampleCalcForce() {

}

func TestCalcForce(t *testing.T) {
	type args struct {
		s1 Star2D
		s2 Star2D
	}
	tests := []struct {
		name string
		args args
		want Vec2
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcForce(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
