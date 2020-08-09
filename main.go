package main

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	src, dist, img, dir string
	loop                int
)

var rootCmd = &cobra.Command{
	Use:   "gen-lgtm",
	Short: "Generate LGTM gif",
	Long:  "Generate LGTM gif (source: https://github.com/8398a7/gen-lgtm)",
	PreRun: func(cmd *cobra.Command, args []string) {
		if src == "" && dir == "" {
			fatal(cmd.Help())
			os.Exit(0)
		}
		if dir != "" {
			dist = ""
		}
		if dir != "" && src != "" {
			fmt.Println("If dir is specified, src cannot be specified.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err  error
			srcs []string
		)
		if dir != "" {
			srcs, err = filepath.Glob(filepath.Join(dir, "*.gif"))
			fatal(err)
		} else {
			srcs = []string{src}
		}

		fn := func(src, dist string) {
			srcFile, lgtm, err := prepare(src, img)
			fatal(err)
			defer srcFile.Close()

			distGif, err := Gen(srcFile, lgtm, loop)
			fatal(err)

			fatal(write(dist, distGif))
		}

		for _, src := range srcs {
			if dist == "" {
				dist = strings.ReplaceAll(src, ".gif", "-dist.gif")
			}
			fn(src, dist)
		}
	},
}

func main() {
	rootCmd.Flags().StringVarP(&src, "src", "s", "", "Source image path")
	rootCmd.Flags().StringVarP(&dist, "dist", "d", "", "Distribution image path")
	rootCmd.Flags().IntVarP(&loop, "loop", "l", -1, "Number of loops in the gif image. If 0 is specified, it is an infinite loop")
	rootCmd.Flags().StringVarP(&img, "image", "i", "", "Overwrite the LGTM image")
	rootCmd.Flags().StringVarP(&dir, "dir", "r", "", "Converts gif images in a specified folder. The image name is \"original-image-dist.gif\"")

	fatal(rootCmd.Execute())
}

func write(dist string, distGif *gif.GIF) error {
	f, err := os.Create(dist)
	if err != nil {
		return fmt.Errorf("failed to create(%s): %w", dist, err)
	}
	defer f.Close()

	if err = gif.EncodeAll(f, distGif); err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}

	return nil
}

func prepare(src, img string) (*os.File, image.Image, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open: %w", err)
	}

	lgtm, err := readLGTM(img)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read lgtm image: %w", err)
	}

	return f, lgtm, nil
}

func fatal(err error) {
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
