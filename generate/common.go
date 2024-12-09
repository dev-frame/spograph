package generate

import (
	"os"
	"path/filepath"

	"github.com/dev-frame/spograph/pkg/gofile"
)

// ListSubDirs lists all subdirectories of the given root directory that have the given subdirectory name.
func ListSubDirs(root string, subDir string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(dirPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && hasSubDir(dirPath, subDir) {
			if subDir == "" {
				dirs = append(dirs, dirPath)
			} else {
				dirs = append(dirs, dirPath+gofile.GetPathDelimiter()+subDir)
			}
		}
		return nil
	})
	return dirs, err
}

func hasSubDir(dirPath string, subDir string) bool {
	_, err := os.Stat(filepath.Join(dirPath, subDir))
	return err == nil || os.IsExist(err)
}
