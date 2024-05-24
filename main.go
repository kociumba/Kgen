package main

import (
	"flag"
	"math"
	"math/rand"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	x := flag.Int("x", 100, "the x dimension of the image")
	y := flag.Int("y", 100, "the y dimension of the image")
	filename := flag.String("fn", "generated", "what the image should be saved as")

	flag.Parse()

	if !flag.Parsed() {
		log.Fatal("fucked up the flags")
	}

	seed := getActualSeed()

	log.Info(seed)
	log.Info("x dimension: ", "x", *x)
	log.Info("y dimension: ", "y", *y)

	canvas := NewCanvas(*x, *y)

	canvas.ColorPixel(seed)

	canvas.Save(*filename)
}

func getActualSeed() int64 {
	ts := int64(time.Now().UnixNano()) / int64(math.Round(math.Pi))
	return ts * int64(rand.Int31())
}
