package main

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	src, dist string
	loop      int
)

var rootCmd = &cobra.Command{
	Use:   "gen-lgtm",
	Short: "Generate LGTM gif",
	Long:  "Generate LGTM gif (source: https://github.com/8398a7/gen-lgtm)",
	PreRun: func(cmd *cobra.Command, args []string) {
		if src == "" {
			fatal(cmd.Help())
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		srcFile, lgtm, err := prepare(src)
		fatal(err)
		defer srcFile.Close()

		distGif, err := Gen(srcFile, lgtm, loop)
		fatal(err)

		fatal(write(dist, distGif))
	},
}

func main() {
	rootCmd.Flags().StringVarP(&src, "src", "s", "", "Source image path.")
	rootCmd.Flags().StringVarP(&dist, "dist", "d", "dist.gif", "Distribution image path.")
	rootCmd.Flags().IntVarP(&loop, "loop", "l", -1, "Number of loops in the gif image. If 0 is specified, it is an infinite loop.")

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

func prepare(src string) (*os.File, image.Image, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open: %w", err)
	}

	lgtm, err := readLGTM()
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
