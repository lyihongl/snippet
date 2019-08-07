package data

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/lyihongl/main/snippet/res"
)

func GetConfig(path string) *map[string]string {
	absPath, _ := filepath.Abs(path)
	//fmt.Println(absPath)
	file, err := os.Open(absPath)
	res.CheckErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string
	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}
	file.Close()
	a := make(map[string]string)
	for _, eachline := range txtLines {
		i := strings.Index(eachline, ":")
		a[eachline[:i]] = eachline[i+1:]
	}
	return &a
}
