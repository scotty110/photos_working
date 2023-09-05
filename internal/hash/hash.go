package hash

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"runtime"
	"sync"

	"fmt"
)

type FileMeta struct {
	FileName string
	FileHash string
}

type hashWorkerInfo struct {
	ctx       context.Context
	inChannel chan string
	fileMap   map[string]FileMeta
	mutex     *sync.Mutex
	wg        *sync.WaitGroup
}

// Funcs
func (data FileMeta) addElement(meta hashWorkerInfo) {
	meta.mutex.Lock()
	defer meta.mutex.Unlock()
	meta.fileMap[data.FileHash] = data
}

func hashFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error") //Need logging for real thing
		return ""
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("error")
	}
	return hex.EncodeToString(hash.Sum(nil)[:])
}

func (meta hashWorkerInfo) hashWorker() {
	defer meta.wg.Done()
	for f := range meta.inChannel {
		h := FileMeta{FileName: f, FileHash: hashFile(f)}
		h.addElement(meta)
	}
}

func HashAll(ctx context.Context, fileList []string) map[string]FileMeta {
	//fileList := file.ReadDir(fileDir)

	// Make Channel
	fileChannel := make(chan string)

	// Make Map for files
	fileMap := make(map[string]FileMeta)

	// MultiProcess Stuff
	ncpu := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	wg.Add(ncpu)
	mu := &sync.Mutex{}

	// Make hashWorkerInfo obj
	meta := hashWorkerInfo{ctx, fileChannel, fileMap, mu, wg}

	for i := 0; i < ncpu; i++ {
		go meta.hashWorker()
	}

	// Dump files to channel
	go func() {
		for i := range fileList {
			fileChannel <- fileList[i]
		}
		close(fileChannel)
	}()

	wg.Wait()
	return fileMap
}

/*
func HashFile(fileName string) []byte {
    //Hash a file

    //Open file for hashing
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //Parameters
    buf := make([]byte, 30*1024)
    sha512 := sha512.New()

    for {
        n, err := file.Read(buf)

        if n > 0 {
            _, err := sha512.Write(buf[:n])
            if err != nil {
                log.Fatal(err)
            }
        }

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Printf("Read %d bytes: %v", n, err)
            break
        }
    }
    sum := sha512.Sum(nil)
    return sum
}
*/
