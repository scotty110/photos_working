package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.a2z.com/consolidate/internal/file"
)

func getExt(f string) string {
	ext := strings.ToLower(filepath.Ext(f))
	return ext
}

func findExt(input []string) map[string]bool {
	extMap := make(map[string]bool)
	for _, v := range input {
		e := getExt(v)
		extMap[e] = true
	}
	return extMap
}

func main() {
	start := time.Now()

	argsLen := len(os.Args[1:])
	if argsLen < 1 || argsLen > 1 {
		fmt.Println("need to make helper")
		return
	}
	// Get Dir from command line
	fileDir := os.Args[len(os.Args)-1]
	fileList := file.ReadDir(fileDir)

	extMap := findExt(fileList)

	for k, _ := range extMap {
		fmt.Println(k)
	}

	fmt.Println("Done")
	fmt.Printf("Duration = %v \n", time.Since(start))
	/*
			   Duration = 16.177036208
		       Duration = 6.833934792s
	*/
}
