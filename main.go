package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/qbart/bmo/bmo"
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"BMO",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		320, 480,
		sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	bkgBMP, _ := sdl.LoadBMP("res/bkg.bmp")
	defer bkgBMP.Free()

	bkgTex, _ := renderer.CreateTextureFromSurface(bkgBMP)
	defer bkgTex.Destroy()

	src := sdl.Rect{0, 0, 320, 480}
	dst := src

	devices := bmo.NewDevices()
	// register by DNS
	devices.RegisterYeeBulb("bmo-yee1")
	devices.RegisterYeeBulb("bmo-yee2")

	screen := bmo.Screen{
		W: 320,
		H: 480,
	}

	components := make([]bmo.IComponent, 0)
	components = append(components, &bmo.Component{
		Rect: sdl.Rect{180, 340, 80, 80},
		Color: bmo.RGB(248, 0, 85),
	})
	// greenButton / rgb(40, 187, 65)
	// aquaButton / rgb(69, 240, 217)
	// yellowButton / rgb(247, 251, 115)
	// display / rgb(211, 255, 219)
	// 40,40, 240, 202

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					p := screen.Position(t)
					for _, c := range components {
						if c.Contains(p) {
							fmt.Println("Clicked {}", p)
						}
					}
				}
			}
		}

		renderer.Clear()
		renderer.Copy(bkgTex, &src, &dst)
		for _, c := range components {
			c.Draw(renderer)
		}
		renderer.Present()
		sdl.Delay(16)
	}
}
