package text

import (
	"github.com/lucasepe/x/image/gg"
)

func DrawGridOverlay(
	dc *gg.Context,
	cols, rows int,
	cellWidth, lineHeight int,
	margin int,
) {
	dc.Push()
	defer dc.Pop()

	// colore griglia: blu semi-trasparente
	dc.SetRGBA(0, 0.4, 1, 0.25)
	dc.SetLineWidth(1)

	// linee verticali
	for c := 0; c <= cols; c++ {
		x := float64(margin + c*cellWidth)
		dc.DrawLine(x, float64(margin), x, float64(margin+rows*lineHeight))
		dc.Stroke()
	}

	// linee orizzontali
	for r := 0; r <= rows; r++ {
		y := float64(margin + r*lineHeight)
		dc.DrawLine(float64(margin), y, float64(margin+cols*cellWidth), y)
		dc.Stroke()
	}
}

func DrawBaselineOverlay(
	dc *gg.Context,
	rows int,
	margin int,
	ascent int,
	lineHeight int,
	width int,
) {
	dc.Push()
	defer dc.Pop()

	// baseline in rosso
	dc.SetRGBA(1, 0, 0, 0.5)
	dc.SetLineWidth(1)

	for r := 0; r < rows; r++ {
		y := margin + r*lineHeight + ascent
		dc.DrawLine(float64(margin), float64(y), float64(width-margin), float64(y))
		dc.Stroke()
	}
}
