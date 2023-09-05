package file

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type metaWorkerInfo struct {
	ctx       context.Context
	inChannel chan string
	wg        *sync.WaitGroup
}

func fixDatetime(fileName string) error {
	srcInfo, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	ctime := readMeta(fileName, srcInfo.ModTime())

	// Set creation time
	err = os.Chtimes(fileName, ctime, ctime)
	if err != nil {
		return err
	}
	return nil
}

func (meta metaWorkerInfo) metaWorker() {
	defer meta.wg.Done()
	for f := range meta.inChannel {
		err := fixDatetime(f)
		if err != nil {
			fmt.Println(f)
		}
	}
}

func FixAll(ctx context.Context, fileList []string) {
	fileChannel := make(chan string)

	// MultiProcess Stuff
	ncpu := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	wg.Add(ncpu)

	// Make hashWorkerInfo obj
	meta := metaWorkerInfo{ctx, fileChannel, wg}

	for i := 0; i < ncpu; i++ {
		go meta.metaWorker()
	}

	// Dump files to channel
	go func() {
		for _, v := range fileList {
			fileChannel <- v
		}
		close(fileChannel)
	}()

	wg.Wait()
}
