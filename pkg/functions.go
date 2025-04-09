package pkg

import "os"

func CleanupFiles(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}
