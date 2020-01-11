package main

import (
	"fmt"
	"log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	sw = 480
	sh = 320
)

func update(screen *ebiten.Image) error {
	return nil
}

func main() {
	if err := ebiten.Run(update, sw, sh, 1, "Proxima"); err != nil {
		log.Fatal(err)
	}
}
