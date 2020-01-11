package main

import (
	"log"
	"fmt"
	"image/color"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	sw = 480
	sh = 320
)

var (
	sampleText      = `The quick brown fox jumps over the lazy dog.`
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	text.Draw(screen, msg, mplusNormalFont, 20, 40, color.White)

	// Draw the sample text
	text.Draw(screen, sampleText, mplusNormalFont, 20, 80, color.White)


	return nil
}

func main() {
	if err := ebiten.Run(update, sw, sh, 1, "Proxima"); err != nil {
		log.Fatal(err)
	}
}
