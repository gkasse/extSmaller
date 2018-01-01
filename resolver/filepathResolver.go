package resolver

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	userHomeDirectoryAlias = "~"
)

// FilepathResolver is resolving full path
type FilepathResolver struct{}

// Resolve is convert to absolute path
// e.g.) '~' convert to user's home directory
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
