package main

import (
	"github.com/gkasse/extSmaller/internal"
	"flag"
	logger "github.com/sirupsen/logrus"
)

var (
	path  = flag.String("d", "", "検索対象のディレクトリ")
	debug = flag.Bool("x", false, "このオプションが指定されている時はデバッグモードで起動する")
)

func main() {
	flag.Parse()
	if *debug {
		logger.SetLevel(logger.DebugLevel)
	}
	logger.Debug("Input value", map[string]string{"d": *path})
	logger.Debug("Input value", map[string]bool{"x": *debug})

	internal.Cli(path)
}
