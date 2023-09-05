package file

import (
	"fmt"
	"time"

	//"log"
	"os"

	"github.com/barasher/go-exiftool"
)

func readMeta(fname string, dTime time.Time) time.Time {

	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return dTime
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(fname)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			if k == "DateCreated" || k == "CreateDate" {
				//fmt.Sprintf("[%v] %v\n", k, v)
				//fmt.Printf("type of time: %T\n", v)
				layout := "2006:01:02 15:04:05"
				rt, err := time.Parse(layout, fmt.Sprintf("%s", v))
				if err != nil {
					return dTime
				}
				return rt
			}
		}
	}
	return dTime
}

/*
func ReadMetaOrig(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		//log.Fatal(err)
		return
	}

	x, err := exif.Decode(f)
	if err != nil {
		//log.Fatal(err)
		return
	} else {
		fmt.Println(x)
	}
	return
}
*/

func fixMeta(origFileName string, newFileName string) error {
	/*
		src, err := os.Open(origFileName)
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
	*/
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
