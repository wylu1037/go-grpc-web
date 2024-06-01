package lattice_manager_grpc

import (
	"path/filepath"
	"runtime"
)

// rootPath is the root directory of this package.
var rootPath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	rootPath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(rootPath, rel)
}
