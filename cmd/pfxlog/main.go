package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.Flags().BoolVarP(&absoluteTime, "absolute", "a", false, "Use absolute time for timestamps")
	rootCmd.Flags().StringVarP(&trimPrefix, "trim", "t", "", "Trim package prefix (ex: github.com/michaelquigley/)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])),
	Short: "pfxlog filter [<file1>...<fileN>]",
	Run:   Filter,
}

var trimPrefix string
var absoluteTime bool

func Filter(_ *cobra.Command, args []string) {
	options := pfxlog.DefaultOptions().SetTrimPrefix(trimPrefix)
	if absoluteTime {
		options = options.SetAbsoluteTime()
	}

	if len(args) > 0 {
		for _, arg := range args {
			f, err := os.OpenFile(arg, os.O_RDONLY, os.ModePerm)
			if err != nil {
				panic(err)
			}
			pfxlog.Filter(f, options)
			if err := f.Close(); err != nil {
				panic(err)
			}
		}
	} else {
		pfxlog.Filter(os.Stdin, options)
	}
}
