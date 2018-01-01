package cmd

import (
	"github.com/gkasse/extSmaller/resolver"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// Cli is core function for pull out of 'large' in filename extension
func Cli(path *string) {
	filepathResolver := resolver.FilepathResolver{}
	directory := filepathResolver.Resolve(*path)
	items, e := ioutil.ReadDir(directory)
	if e != nil {
		logger.Error("Error has occurred!", e)
		os.Exit(1)
	}
	logger.Debug("Fetched directory's files counts", len(items))

	if len(items) <= 0 {
		logger.Info("Could not found files in target directory. End of action.")
		os.Exit(0)
	}

	extensionResolver := resolver.ExtensionResolver{}
	for _, fileInfo := range items {
		filename := fileInfo.Name()
		logger.Info("Next file... ", filename)
		available, e := extensionResolver.Available(directory, filename)
		if e != nil {
			logger.Error("Error occurred. ", e)
			continue
		}
		if !available {
			logger.Info("This file is not target.")
			continue
		}

		logger.Debug("Start extension resolving. filename: ", filename)
		e = extensionResolver.Resolve(directory, filename)
		if e != nil {
			logger.Error("Could not resolving extension.", e)
		}
		logger.Info("Complete. filename before resolve: ", filename)
	}
}
