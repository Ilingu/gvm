package gvm

import (
	"gvm-windows/gvm/utils"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func ScanAndDelete(limitDate int64) (int, error) {
	appDir, err := utils.GenerateAppDataPath()
	if err != nil {
		return 0, err
	}

	files, err := os.ReadDir(appDir)
	if err != nil {
		return 0, err
	}

	filesDeleted := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileInfo, err := file.Info()
		if err != nil {
			continue
		}

		if fileInfo.ModTime().UnixMilli() <= limitDate {
			wg.Add(1)
			filesDeleted++
			go func() {
				filepath := filepath.Join(appDir, fileInfo.Name())
				err := os.Remove(filepath)
				if err != nil {
					filesDeleted--
				}

				wg.Done()
			}()
		}
	}
	wg.Wait()

	return filesDeleted, nil
}
