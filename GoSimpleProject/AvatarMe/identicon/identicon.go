package identicon

import (
	"crypto/md5"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type gridPoint struct {
	value byte
	index int
}

type avatar struct {
	hash       [16]byte
	color      [3]byte
	grid       []byte
	gridPoints []gridPoint
	pixelMap   []image.Rectangle
}

func Generate(name string) *avatar {
	a := avatar{}
	a.hashName(name)
	a.pickColor()
	a.buildGrid()
	a.filterOddSquares()
	a.buildPixelMap()
	return &a
}

func (a avatar) WriteImage(name string) error {
	var img = image.NewRGBA(image.Rect(0, 0, 250, 250))
	//col := color.RGBA{R: i.color[0], G: i.color[1], B: i.color[2], A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 0xff}}, image.Point{}, draw.Src)
	for _, pixel := range a.pixelMap {
		a.rect(img, pixel.Min.X, pixel.Min.Y, pixel.Max.X, pixel.Max.Y)
	}
	myFile, _ := os.Create(name + ".png")
	return png.Encode(myFile, img)
}

func (a *avatar) hashName(s string) {
	a.hash = md5.Sum([]byte(s))
}

func (a *avatar) pickColor() {
	rgb := [3]byte{}
	// next we copy the first 3 values from the hash to the rgb array
	copy(rgb[:], a.hash[:3])
	// we than assign it to the color value
	a.color = rgb
}

func (a *avatar) buildGrid() {
	var grid []byte
	for i := 0; i < len(a.hash) && i+3 <= len(a.hash)-1; i += 3 {
		chunk := make([]byte, 5)
		copy(chunk, a.hash[i:i+3])
		chunk[3] = chunk[1]           // mirror the second value in the chunk
		chunk[4] = chunk[0]           // mirror the first value in the chunk
		grid = append(grid, chunk...) // append the chunk to the grid
	}
	a.grid = grid // set the grid property on the identity
}

func (a *avatar) filterOddSquares() {
	var grid []gridPoint          // create a placeholder to hold the values of the loop
	for i, code := range a.grid { // loop over the grid
		if code%2 == 0 { // check if the value is odd or not
			// create a new Grid point where we save the value and the index in the grid
			point := gridPoint{
				value: code,
				index: i,
			}
			// append the item to the new grid
			grid = append(grid, point)
		}
	}
	// set the property
	a.gridPoints = grid
}

func (a *avatar) buildPixelMap() {
	var drawingPoints []image.Rectangle // define placeholder for drawing points

	// Closure, this function returns a drawing point
	pixelFunc := func(p gridPoint) image.Rectangle {
		// This is the formula, we use the index from the grid point to calculate the horizontal dimension
		horizontal := (p.index % 5) * 50
		// This is the formula, we use the index from the grid point to calculate the vertical dimension
		vertical := (p.index / 5) * 50
		// this is the top left point with x and the y
		topLeft := image.Point{X: horizontal, Y: vertical}
		// the bottom right point is just the top left point +50 because 1 block in the grid is 50x50
		bottomRight := image.Point{X: horizontal + 50, Y: vertical + 50}

		return image.Rectangle{ // We then return the drawing point
			Min: topLeft,
			Max: bottomRight,
		}
	}

	for _, gridPoint := range a.gridPoints {
		// for every gridPoint we calculate the drawing points, and we add them to the array
		drawingPoints = append(drawingPoints, pixelFunc(gridPoint))
	}
	a.pixelMap = drawingPoints // set the drawing point value on the identity
}

func (a avatar) rect(img *image.RGBA, x1, x2, x3, x4 int) {
	redRect := image.Rect(x1, x2, x3, x4) //  geometry of 2nd rectangle which we draw atop above rectangle
	col := color.RGBA{R: a.color[0], G: a.color[1], B: a.color[2], A: 255}
	draw.Draw(img, redRect, &image.Uniform{C: col}, image.Point{}, draw.Src)
}
