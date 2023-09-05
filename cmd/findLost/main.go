package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.a2z.com/consolidate/internal/file"
	"golang.a2z.com/consolidate/internal/hash"
	"golang.a2z.com/consolidate/internal/sets"
)

func makeMap(input []string) map[string]string {
	hashMap := make(map[string]string)
	for _, v := range input {
		fname := filepath.Base(v)
		hash := strings.Split(fname, ".")[0]
		hashMap[hash] = v
	}
	return hashMap
}

func main() {
	start := time.Now()
	ctx := context.Background()

	argsLen := len(os.Args[1:])
	if argsLen < 2 || argsLen > 2 {
		fmt.Println("need to make helper")
		return
	}
	// Get Dir from command line
	fileDirAlpha := os.Args[len(os.Args)-2]
	fileListAlpha := file.ReadDir(fileDirAlpha)

	fileDirBeta := os.Args[len(os.Args)-1]
	fileListBeta := file.ReadDir(fileDirBeta)

	// Create map of (hash -> file)
	//fileMap := hash.HashAll(ctx, fileList)

	// Make dir of unique files
	//file.CopyAll(ctx, fileMap, fileDir)

	alphaMap := hash.HashAll(ctx, fileListAlpha)
	betaMap := hash.HashAll(ctx, fileListBeta)

	moveMap := sets.DiffCopy(betaMap, alphaMap)

	// Move
	file.CopyAll(ctx, moveMap, fileDirBeta)

	fmt.Println("Done")
	fmt.Printf("Duration = %v \n", time.Since(start))
	/*
			   Duration = 16.177036208
		       Duration = 6.833934792s
	*/
}
