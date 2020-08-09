package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	version   = "master"
	commit    = "master"
	date      = time.Now().String()
	goversion = "-"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			v := map[string]interface{}{
				"Version":   version,
				"Commit":    commit,
				"BuildDate": date,
				"GoVersion": goversion,
			}
			b, err := json.Marshal(v)
			fatal(err)

			fmt.Println(string(b))
		},
	}

	rootCmd.AddCommand(versionCmd)
}
