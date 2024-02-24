package utils

import (
	"gitlet/config"
	"log"
	"os"
	"path/filepath"
)

type where int

const (
	BLOB where = iota
	ADDSTAGE
	RMSTAGE
)

func WriteFile(filename string, data string) {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFileBytes(filename string, bytes []byte) {
	err := os.WriteFile(filename, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

/* Move files from source to dest. */
func MoveFiles(source where, dest where) {
	sourcePath := getWhere(source)
	destPath := getWhere(dest)

	files, err := os.ReadDir(sourcePath)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        oldPath := filepath.Join(sourcePath, file.Name())
        newPath := filepath.Join(destPath, file.Name())

        err := os.Rename(oldPath, newPath)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func getWhere(w where) string {
	switch w {
	case BLOB:
		return config.BLOB
	case ADDSTAGE:
		return config.ADDSTAGE
	case RMSTAGE:
		return config.RMSTAGE
	default:
		return ""
	}
}

/* Remove all file in dest. */
func RemoveFiles(dest where) {
	destPath := getWhere(dest)
	files, err := os.ReadDir(destPath)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        filePath := filepath.Join(destPath, file.Name())
        err := os.Remove(filePath)
        if err != nil {
            log.Fatal(err)
        }
    }
}
