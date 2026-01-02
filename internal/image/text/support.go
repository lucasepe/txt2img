package text

import (
	"math"

	"github.com/lucasepe/x/image/gg"
	"golang.org/x/image/font"
)

type AdvanceModel struct {
	Base  float64
	Cache map[rune]int
	DC    *gg.Context
}

func NewAdvanceModel(face font.Face) *AdvanceModel {
	dc := gg.NewContext(100, 100)
	dc.SetFontFace(face)

	// base = advance dello spazio
	base, _ := dc.MeasureString(" ")

	if base <= 0 {
		base = 1
	}

	return &AdvanceModel{
		Base:  base,
		Cache: make(map[rune]int),
		DC:    dc,
	}
}

func (m *AdvanceModel) Cols(r rune) int {
	if v, ok := m.Cache[r]; ok {
		return v
	}

	w, _ := m.DC.MeasureString(string(r))
	cols := int(math.Round(w / m.Base))

	if cols < 1 {
		cols = 1
	}

	m.Cache[r] = cols
	return cols
}

func (m *AdvanceModel) LineCols(line string) int {
	total := 0
	for _, r := range line {
		total += m.Cols(r)
	}
	return total
}

func MeasureGridAdvance(lines []string, adv *AdvanceModel) (cols, rows int) {
	max := 0
	for _, line := range lines {
		n := adv.LineCols(line)
		if n > max {
			max = n
		}
	}
	return max, len(lines)
}

func DrawAdvanceText(
	dc *gg.Context,
	lines []string,
	adv *AdvanceModel,
	margin int,
	lineHeight int,
	ascent int,
) {
	y := margin + ascent

	for _, line := range lines {
		x := margin
		for _, r := range line {
			dc.DrawString(string(r), float64(x), float64(y))
			w, _ := adv.DC.MeasureString(string(r))
			x += int(math.Round(w))
		}
		y += lineHeight
	}
}
