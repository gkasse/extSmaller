package resolver

import (
	"os"
	"strings"
	"path/filepath"
	"os/user"
	logger "github.com/sirupsen/logrus"
)

const (
	userHomeDirectoryAlias = "~"
)

type FilepathResolver struct{}

func (r FilepathResolver) Resolve(path string) string {
	currentUser, e := user.Current()
	if e != nil {
		logger.Error("Error has occurred!", e)
		os.Exit(1)
	}

	logger.Debug("Current user's home directory: ", currentUser.HomeDir)
	resolvedPath := strings.Replace(path, userHomeDirectoryAlias, currentUser.HomeDir, -1)
	logger.Debug("Replace complete. result: ", resolvedPath)
	return filepath.Clean(resolvedPath)
}
