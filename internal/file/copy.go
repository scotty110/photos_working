package file

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	//"flag"

	"golang.a2z.com/consolidate/internal/hash"
)

// Structs
/*
type FileMeta struct {
	FileName string
	FileHash string
}
*/

type copyWorkerInfo struct {
	ctx       context.Context
	inChannel chan hash.FileMeta
	newDir    string
	wg        *sync.WaitGroup
}

func setDatetime(origFileName string, newFileName string) error {
	srcInfo, err := os.Stat(origFileName)
	if err != nil {
		return err
	}

	err = os.Chmod(newFileName, srcInfo.Mode())
	if err != nil {
		return err
	}

	ctime := readMeta(origFileName, srcInfo.ModTime())

	// Set creation time
	err = os.Chtimes(newFileName, ctime, ctime)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(data hash.FileMeta, newDir string) error {
	// Create new file
	ext := strings.ToLower(filepath.Ext(data.FileName))
	/*
		absFile, err := filepath.Abs(data.FileName)
		if err != nil {
			return err
		}
		dir := filepath.Dir(absFile)
	*/

	newFileName := fmt.Sprintf("%s/%s%s", newDir, data.FileHash, ext)

	// Open File
	src, err := os.Open(data.FileName)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	setDatetime(data.FileName, newFileName)
	return nil
}

func (meta copyWorkerInfo) copyWorker() {
	defer meta.wg.Done()
	for f := range meta.inChannel {
		err := copyFile(f, meta.newDir)
		if err != nil {
			fmt.Println(f.FileName)
		}
	}
}

func makeDir(fileDir string) string {
	newDir := fmt.Sprintf("%s_cleaned", fileDir)
	err := os.Mkdir(newDir, 0755)
	if err != nil {
		fmt.Println("Something")
	}
	return newDir
}

func CopyAll(ctx context.Context, fileMap map[string]hash.FileMeta, origDir string) {
	newDir := makeDir(origDir)
	hashChannel := make(chan hash.FileMeta)

	// MultiProcess Stuff
	ncpu := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	wg.Add(ncpu)

	// Make hashWorkerInfo obj
	meta := copyWorkerInfo{ctx, hashChannel, newDir, wg}

	for i := 0; i < ncpu; i++ {
		go meta.copyWorker()
	}

	// Dump files to channel
	go func() {
		for _, v := range fileMap {
			hashChannel <- v
		}
		close(hashChannel)
	}()

	wg.Wait()
}
