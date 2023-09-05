package file

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"golang.a2z.com/consolidate/internal/hash"
)

type pathWorkerInfo struct {
	ctx       context.Context
	inChannel chan hash.FileMeta
	wg        *sync.WaitGroup
}

func fixPath(data hash.FileMeta) error {
	ext := strings.ToLower(filepath.Ext(data.FileName))
	absFile, err := filepath.Abs(data.FileName)
	if err != nil {
		return err
	}
	dir := filepath.Dir(absFile)

	newFileName := fmt.Sprintf("%s/%s%s", dir, data.FileHash, ext)

	err = os.Rename(data.FileName, newFileName)
	if err != nil {
		//fmt.Println(newFileName)
		return err
	}
	return nil
}

func (meta pathWorkerInfo) pathWorker() {
	defer meta.wg.Done()
	for f := range meta.inChannel {
		err := fixPath(f)
		if err != nil {
			fmt.Println(err)
			fmt.Println(f)
		}
	}
}

func PathFix(ctx context.Context, fileHash map[string]hash.FileMeta) {
	fileChannel := make(chan hash.FileMeta)

	// MultiProcess Stuff
	ncpu := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	wg.Add(ncpu)

	// Make hashWorkerInfo obj
	meta := pathWorkerInfo{ctx, fileChannel, wg}

	for i := 0; i < ncpu; i++ {
		go meta.pathWorker()
	}

	// Dump files to channel
	go func() {
		for _, v := range fileHash {
			fileChannel <- v
		}
		close(fileChannel)
	}()

	wg.Wait()
}
