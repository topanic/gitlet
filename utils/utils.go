package utils

import (
	"crypto/sha1"
	"fmt"
	"gitlet/config"
	"log"
	"os"
	"path/filepath"
	"time"
)

type where int

const (
	BLOB where = iota
	ADDSTAGE
	RMSTAGE
)

func GetArgsNum(args []string) int {
	return len(args)
}

/* generate a SHA-1 code */
func GenerateID() string {
	timestamp := time.Now().String()

	hasher := sha1.New()
	hasher.Write([]byte(timestamp))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash)
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
