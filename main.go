package main

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	// "io/ioutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"log"
)

const (
	sw      = 480
	sh      = 320
	name    = "Proxima"
	version = "0.1"
)

var (
	debug                  string
	fontpixel              font.Face
	smallfontpixel         font.Face
	buttons                []*Button
	prevMouseClickDuration int
)

type Button struct {
	t    string
	x    float64
	y    float64
	size float64
	c    color.Color
}

func (b *Button) contains(x, y int) bool {
	return float64(x) >= b.x &&
		float64(x) <= b.x+b.size-1 &&
		float64(y) >= b.y &&
		float64(y) <= b.y+b.size-1
}

func init() {
	// b, err := ioutil.ReadFile("custom font path")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	b := fonts.MPlus1pRegular_ttf

	f, err := truetype.Parse(b)
	if err != nil {
		log.Println(err)
		return
	}

	fontpixel = truetype.NewFace(f, &truetype.Options{
		Size:    18,
		Hinting: font.HintingFull,
	})
	smallfontpixel = truetype.NewFace(f, &truetype.Options{
		Size:    12,
		Hinting: font.HintingFull,
	})

	buttons = make([]*Button, 0)
	buttons = append(buttons, &Button{
		t:    "Blue",
		x:    60,
		y:    70,
		size: 100,
		c:    color.RGBA{0x00, 0x00, 0xff, 0xff},
	})
	buttons = append(buttons, &Button{
		t:    "Green",
		x:    190,
		y:    70,
		size: 100,
		c:    color.RGBA{0x00, 0xff, 0x00, 0xff},
	})
	buttons = append(buttons, &Button{
		t:    "White",
		x:    60,
		y:    200,
		size: 100,
		c:    color.RGBA{0xff, 0xff, 0xff, 0xff},
	})
	buttons = append(buttons, &Button{
		t:    "Off",
		x:    190,
		y:    200,
		size: 100,
		c:    color.RGBA{0x00, 0x00, 0x00, 0xff},
	})
}

func update(screen *ebiten.Image) error {
	mx, my := ebiten.CursorPosition()

	mouseClickDuration := inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft)
	if mouseClickDuration == 0 && prevMouseClickDuration > 0 {
		for _, button := range buttons {
			if button.contains(mx, my) {
				debug = button.t
				break
			}
		}
	}
	prevMouseClickDuration = mouseClickDuration

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{0x57, 0x38, 0x5d, 0xff})

	ebitenutil.DrawRect(screen, 0, 20, 120, 24, color.RGBA{0xd2, 0x80, 0x7e, 0xff})
	text.Draw(screen, name, fontpixel, 20, 38, color.RGBA{0x56, 0x04, 0x19, 0xff})

	for _, button := range buttons {
		ebitenutil.DrawRect(screen, button.x, button.y, button.size, button.size, button.c)
	}

	msg := fmt.Sprintf("v%s", version)
	text.Draw(screen, msg, smallfontpixel, 450, 310, color.RGBA{0xce, 0xa9, 0x9e, 0xff})

	ebitenutil.DebugPrint(screen, debug)
	return nil
}

func main() {
	if err := ebiten.Run(update, sw, sh, 1, name); err != nil {
		log.Fatal(err)
	}
}
