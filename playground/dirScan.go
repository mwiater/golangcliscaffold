package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type FileInfo struct {
	Name      string
	Size      int64
	IsDir     bool
	FileCount int
}

func getLargestFilesAndDirectories(path string, n int) ([]FileInfo, []FileInfo, error) {
	var files []FileInfo
	var dirs []FileInfo

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, FileInfo{
				Name:      filePath,
				Size:      info.Size(),
				IsDir:     false,
				FileCount: 0,
			})
		} else {
			dirSize, err := calculateDirectorySize(filePath)
			if err != nil {
				return err
			}
			dirs = append(dirs, FileInfo{
				Name:      filePath,
				Size:      dirSize,
				IsDir:     true,
				FileCount: 0,
			})
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})

	if len(files) > n {
		files = files[:n]
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Size > dirs[j].Size
	})

	if len(dirs) > n {
		dirs = dirs[:n]
	}

	for i := range dirs {
		fileCount, err := countFiles(dirs[i].Name)
		if err != nil {
			return nil, nil, err
		}

		dirs[i].FileCount = fileCount
	}

	return files, dirs, nil
}

func calculateDirectorySize(dirPath string) (int64, error) {
	var size int64

	err := filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return size, nil
}

func countFiles(path string) (int, error) {
	fileCount := 0

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return fileCount, nil
}

func main() {
	path := "/home/matt/projects"
	n := 10

	files, dirs, err := getLargestFilesAndDirectories(path, n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Largest Files:", len(files))
	for _, file := range files {
		fmt.Printf("Name: %s, Size: %d\n", file.Name, file.Size)
	}

	fmt.Println("\nLargest Directories:", len(dirs))
	for _, dir := range dirs {
		fmt.Printf("Name: %s, Size: %d, File Count: %d\n", dir.Name, dir.Size, dir.FileCount)
	}
}
