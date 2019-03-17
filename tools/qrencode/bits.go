package qrencode

import (
	"bytes"
	"image"
	"image/color"
)

/*BitVector*/
type BitVector struct {
	boolBitVector
}

func (v *BitVector) AppendBits(b BitVector) {
	v.boolBitVector.AppendBits(b.boolBitVector)
}

func (v *BitVector) String() string {
	b := bytes.Buffer{}
	for i, l := 0, v.Length(); i < l; i++ {
		if v.Get(i) {
			b.WriteString("X")
		} else {
			b.WriteString(".")
		}
	}
	return b.String()
}

/*BitGrid*/
type BitGrid struct {
	boolBitGrid
}

func NewBitGrid(width, height int) *BitGrid {
	return &BitGrid{newBoolBitGrid(width, height)}
}

func (g *BitGrid) String() string {
	b := bytes.Buffer{}
	for y, w, h := 0, g.Width(), g.Height(); y < h; y++ {
		for x := 0; x < w; x++ {
			if g.Empty(x, y) {
				b.WriteString(" ")
			} else if g.Get(x, y) {
				b.WriteString("1")
			} else {
				b.WriteString("0")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

// Return an image of the grid, with black blocks for true items and
// white blocks for false items, with the given block size and a
// default margin.
func (g *BitGrid) Image(blockSize int) image.Image {
	return g.ImageWithMargin(blockSize, 4)
}

// Return an image of the grid, with black blocks for true items and
// white blocks for false items, with the given block size and margin.
func (g *BitGrid) ImageWithMargin(blockSize, margin int) image.Image {
	width := blockSize * (2*margin + g.Width())
	height := blockSize * (2*margin + g.Height())

	i := image.NewGray16(image.Rect(0, 0, width, height))
	for y := 0; y < blockSize*margin; y++ {
		for x := 0; x < width; x++ {
			i.Set(x, y, color.White)
			i.Set(x, height-1-y, color.White)
		}
	}
	for y := blockSize * margin; y < height-blockSize*margin; y++ {
		for x := 0; x < blockSize*margin; x++ {
			i.Set(x, y, color.White)
			i.Set(width-1-x, y, color.White)
		}
	}
	for y, w, h := 0, g.Width(), g.Height(); y < h; y++ {
		for x := 0; x < w; x++ {
			x0 := blockSize * (x + margin)
			y0 := blockSize * (y + margin)
			c := color.White
			if g.Get(x, y) {
				c = color.Black
			}
			for dy := 0; dy < blockSize; dy++ {
				for dx := 0; dx < blockSize; dx++ {
					i.Set(x0+dx, y0+dy, c)
				}
			}
		}
	}
	return i
}

/*boolBitVector*/
type boolBitVector struct {
	bits []bool
}

func (v *boolBitVector) Length() int {
	return len(v.bits)
}

func (v *boolBitVector) Get(i int) bool {
	return v.bits[i]
}

func (v *boolBitVector) AppendBit(b bool) {
	v.bits = append(v.bits, b)
}

func (v *boolBitVector) Append(b, count int) {
	for i := uint(count); i > 0; i-- {
		v.AppendBit((b>>(i-1))&1 == 1)
	}
}

func (v *boolBitVector) AppendBits(b boolBitVector) {
	v.bits = append(v.bits, b.bits...)
}

/*boolBitGrid*/
type boolBitGrid struct {
	width, height int
	bits          []bool
}

func newBoolBitGrid(width, height int) boolBitGrid {
	return boolBitGrid{width, height, make([]bool, 2*width*height)}
}

func (g *boolBitGrid) Width() int {
	return g.width
}

func (g *boolBitGrid) Height() int {
	return g.height
}

func (g *boolBitGrid) Empty(x, y int) bool {
	return !g.bits[2*(x+y*g.width)]
}

func (g *boolBitGrid) Get(x, y int) bool {
	return g.bits[2*(x+y*g.width)+1]
}

func (g *boolBitGrid) Set(x, y int, v bool) {
	g.bits[2*(x+y*g.width)] = true
	g.bits[2*(x+y*g.width)+1] = v
}

func (g *boolBitGrid) Clear() {
	for i, _ := range g.bits {
		g.bits[i] = false
	}
}
