package text

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/lucasepe/x/image/bdf"
	"github.com/lucasepe/x/image/bdf/fonts"
	"github.com/lucasepe/x/image/gg"
)

// RenderOptions defines the configuration options for TextToImage.
//
// Each field controls a specific aspect of text rendering, image sizing
// and layout. Unless otherwise specified, fields default to 0 (or false)
// and are overridden by internal defaults of the renderer.
type RenderOptions struct {
	// Margin sets the margin (padding) around the text block, in pixels.
	// The margin is applied equally on all sides.
	Margin int

	// LineSpacing multiplies the base line height (ascent+descent).
	// Values <= 0 default to 1.0.
	// This is a grid-based spacing, not typographic leading.
	LineSpacing float64

	// TransparentBackground, when true, makes the background transparent
	// instead of solid white.
	TransparentBackground bool

	TextColor color.Color

	BackgroundColor color.Color

	DebugBaseline bool
	DebugGrid     bool
}

func RenderGG(text string, opts RenderOptions) (*gg.Context, error) {
	fnt, err := bdf.LoadFont(fonts.Cozette2X())
	if err != nil {
		return nil, fmt.Errorf("unable to load fonts: %s", err)
	}

	face := fnt.NewFace()

	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " \t")
	}

	adv := NewAdvanceModel(face)

	// misura griglia in "colonne logiche"
	cols, rows := MeasureGridAdvance(lines, adv)

	metrics := face.Metrics()
	ascent := metrics.Ascent.Ceil()
	descent := metrics.Descent.Ceil()

	baseLineHeight := ascent + descent

	// LineSpacing reinterpretato come moltiplicatore discreto
	lineSpacing := opts.LineSpacing
	if lineSpacing <= 0 {
		lineSpacing = 1.0
	}

	effectiveLineHeight := int(float64(baseLineHeight) * lineSpacing)

	// larghezza basata sull'advance
	cellWidth := int(math.Round(adv.Base))

	width := cols*cellWidth + 2*opts.Margin
	height := rows*effectiveLineHeight + 2*opts.Margin

	dc := gg.NewContext(width, height)
	dc.SetFontFace(face)

	// background
	if opts.BackgroundColor != nil {
		dc.SetColor(opts.BackgroundColor)
		dc.Clear()
	}

	// text color
	if opts.TextColor != nil {
		dc.SetColor(opts.TextColor)
	} else {
		dc.SetRGB(0, 0, 0)
	}

	DrawAdvanceText(
		dc,
		lines,
		adv,
		opts.Margin,
		effectiveLineHeight,
		ascent,
	)

	if opts.DebugGrid {
		DrawGridOverlay(
			dc,
			cols,
			rows,
			cellWidth,
			effectiveLineHeight,
			opts.Margin,
		)
	}

	if opts.DebugBaseline {
		DrawBaselineOverlay(
			dc,
			rows,
			opts.Margin,
			ascent,
			effectiveLineHeight,
			width,
		)
	}

	return dc, nil
}
