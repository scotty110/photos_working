package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.a2z.com/consolidate/internal/file"
	"golang.a2z.com/consolidate/internal/hash"
)

func main() {
	start := time.Now()
	ctx := context.Background()

	argsLen := len(os.Args[1:])
	if argsLen < 1 || argsLen > 1 {
		fmt.Println("need to make helper")
		return
	}
	// Get Dir from command line
	fileDir := os.Args[len(os.Args)-1]
	fileList := file.ReadDir(fileDir)

	fileHash := hash.HashAll(ctx, fileList)
	file.PathFix(ctx, fileHash)

	fmt.Println("Done")
	fmt.Printf("Duration = %v \n", time.Since(start))
}
