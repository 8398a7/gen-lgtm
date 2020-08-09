package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"os"
)

func Gen(src io.Reader, lgtm image.Image, loop int) (*gif.GIF, error) {
	srcGif, err := gif.DecodeAll(src)
	if err != nil {
		return nil, fmt.Errorf("failed to decode to gif: %w", err)
	}

	cfg := srcGif.Config
	bounds := lgtm.Bounds()
	if cfg.Width < bounds.Dx() || cfg.Height < bounds.Dy() {
		return nil, fmt.Errorf("The LGTM image is larger than the source gif image.(src: %dx%d, lgtm: %dx%d\n",
			cfg.Width, cfg.Height,
			bounds.Dx(), bounds.Dy())
	}

	out, err := composite(srcGif, lgtm, loop)
	if err != nil {
		return nil, fmt.Errorf("failed to composite: %w", err)
	}

	return out, nil
}

func composite(src *gif.GIF, lgtm image.Image, loop int) (*gif.GIF, error) {
	cfg := src.Config
	dist := &gif.GIF{
		Image:     make([]*image.Paletted, 0, len(src.Image)),
		Delay:     src.Delay,
		LoopCount: loop,
		Disposal:  src.Disposal,
		Config:    cfg,
	}

	x := cfg.Width/2 - lgtm.Bounds().Dx()/2
	y := cfg.Height/2 - lgtm.Bounds().Dy()/2

	for _, frame := range src.Image {
		draw.Draw(frame, frame.Bounds(), lgtm, image.Point{X: -x, Y: -y}, draw.Over)
		dist.Image = append(dist.Image, frame)
	}

	return dist, nil
}

func readLGTM(img string) (image.Image, error) {
	var r io.Reader
	if img != "" {
		f, err := os.Open(img)
		if err != nil {
			return nil, fmt.Errorf("failed to open: %w", err)
		}
		defer f.Close()
		r = f
	} else {
		f, err := Assets.Open("/assets/lgtm.png")
		if err != nil {
			return nil, fmt.Errorf("failed to open assets: %w", err)
		}
		defer f.Close()
		r = f
	}

	out, err := png.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("failed to decode to png: %w", err)
	}
	return out, nil
}
