package dirs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/mwiater/golangcliscaffold/common"
	"github.com/spf13/viper"
)

type Dir struct {
	Path            string
	Depth           int
	BytesSize       int64
	PrettyBytesSize string
}

var dirsVisited []string
var dirs []Dir

func ReadDirDepth(dirPath string) ([]Dir, error) {
	currentDir := strings.ReplaceAll(dirPath, viper.GetString("path"), "")
	currentDepth := len(strings.Split(currentDir, "/"))

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			_, err := ReadDirDepth(filepath.Join(dirPath, file.Name()))
			if err != nil {
				return nil, err
			}
			if viper.GetInt("depth") >= currentDepth {

				if !common.SliceContains(dirsVisited, filepath.Join(dirPath, file.Name())) {
					dirSize, _ := DirSizeBytes(filepath.Join(dirPath, file.Name()))
					if dirSize > int64(viper.GetInt("mindirsize")*1000000) {
						dir := Dir{}
						dir.Path = filepath.Join(dirPath, file.Name())
						dir.Depth = viper.GetInt("depth")
						dir.BytesSize = dirSize
						dir.PrettyBytesSize = common.PrettyBytes(dirSize)

						dirs = append(dirs, dir)
					}

				}
			}
		}
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].BytesSize > dirs[j].BytesSize
	})

	return dirs, nil
}

// DirSize returns the size of a directory in bytes.
func DirSizeBytes(dirPath string) (int64, error) {
	var size int64
	fsi, err := os.Open(dirPath)
	if err != nil {
		return 0, err
	}
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	defer fsi.Close()
	return size, nil
}

func PrintResults(dirs []Dir) {
	fmt.Println()
	common.PrintColor("forestgreen", "background", fmt.Sprintf("Largest directories in: %s", viper.GetString("path")))
	fmt.Println("----------------------------------------------------------")
	fmt.Println()

	spacing := make(map[string]int)
	highWaterMark := 0

	for _, dir := range dirs {
		if len(dir.Path) > highWaterMark {
			highWaterMark = len(dir.Path)
		}

		spacing[dir.Path] = len(dir.Path)
	}

	for _, dir := range dirs {
		padding := strconv.Itoa(highWaterMark + 2)
		if dir.BytesSize >= int64(viper.GetInt("highlight")*1000000) {
			color.HEXStyle("000", common.AllHex["yellow1"]).Printf("%-"+padding+"s %10s\n", dir.Path, dir.PrettyBytesSize)
		} else {
			color.HEXStyle(common.AllHex["steelblue2"]).Printf("%-"+padding+"s %10s\n", dir.Path, dir.PrettyBytesSize)
		}
	}
}
