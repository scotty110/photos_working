package file

import (
    "io/fs"
    "path/filepath"
    "log"
    "os"
)
/*
func ReadDir(rootDir string) []string {
    var a []string
    filepath.WalkDir(rootDir, func(s string, d fs.DirEntry, e error) error {
        if e != nil {
            log.Fatal(e)
            return e
        }
        if filepath.Ext(d.Name()) == ".jpg" {
            a = append(a, s)
        }
        return nil
    })
    return a
}
*/
func ReadDir(rootDir string) []string {
    var returnFiles []string
    fs.WalkDir( os.DirFS(rootDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
            return nil //Will eventually just want to skip
		}
        fullPath := filepath.Join(rootDir, path)
        if stat,err := os.Stat(fullPath); err == nil && !stat.IsDir() {
            if filepath.Ext(fullPath) != ".DS_Store" {
                returnFiles = append(returnFiles, fullPath)
            }
        }

        return nil
    })
    return returnFiles
}
