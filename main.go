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

	screen := bmo.Screen{
		Rect: sdl.Rect{0, 0, 320, 480},
	}

	window, err := sdl.CreateWindow(
		"BMO",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screen.Rect.W, screen.Rect.H,
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

	devices := bmo.NewDevices()
	// register by DNS
	devices.RegisterYeeBulb("bmo-yee1")
	devices.RegisterYeeBulb("bmo-yee2")

	components := make([]bmo.IComponent, 0)
	components = append(components, &bmo.Component{
		Rect: sdl.Rect{180, 340, 80, 80},
		Color: bmo.RGB(248, 0, 85),
	})
	components = append(components, &bmo.Component{
		Rect: sdl.Rect{40, 40, 240, 202},
		Color: bmo.RGB(211, 255, 219),
	})
	components[0].Show(true)
	components[0].OnMousePressed(func(event bmo.MouseEvent) {
		components[1].Show(true)
		fmt.Println("Point {}", event.P)
	})
	// greenButton / rgb(40, 187, 65)
	// aquaButton / rgb(69, 240, 217)
	// yellowButton / rgb(247, 251, 115)

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
							c.TriggerOnMousePressed(p)
						}
					}
				}
			}
		}

		renderer.Clear()
		renderer.Copy(bkgTex, &screen.Rect, &screen.Rect)
		for _, c := range components {
			c.Draw(renderer)
		}
		renderer.Present()
		sdl.Delay(16)
	}

	fmt.Println("")
}
