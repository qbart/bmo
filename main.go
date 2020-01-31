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

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.MouseButtonEvent:
				if t.State == sdl.RELEASED {
					fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)

				}
			default:
				fmt.Println("event {}", t)
			}
		}

		renderer.Clear()
		renderer.Copy(bkgTex, &src, &dst)
		renderer.Present()
		sdl.Delay(16)
	}
}
