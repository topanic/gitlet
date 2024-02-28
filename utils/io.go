package utils

import (
	"gitlet/config"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Where int

const (
	BLOB Where = iota
	ADDSTAGE
	RMSTAGE
	// WORK
)

func WriteFile(filePath string, data string) {
	err := os.WriteFile(filePath, []byte(data), 0644)
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

func GetWhere(w Where) string {
	switch w {
	case BLOB:
		return config.BLOB
	case ADDSTAGE:
		return config.ADDSTAGE
	case RMSTAGE:
		return config.RMSTAGE
	// case WORK:
	// 	return config.WORKTREE
	default:
		return ""
	}
}

/* Move files from source to dest. */
func MoveFiles(source Where, dest Where) {
	sourcePath := GetWhere(source)
	destPath := GetWhere(dest)

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

/* Move target file in source to dest. */
func MoveFile(source Where, dest Where, filename string) {
	sourcePath := GetWhere(source)
	destPath := GetWhere(dest)

	files, err := os.ReadDir(sourcePath)
    if err != nil {
        log.Fatal(err)
    }

	for _, file := range files {
		if file.Name() == filename {
			oldPath := filepath.Join(sourcePath, file.Name())
			newPath := filepath.Join(destPath, file.Name())
	
			err := os.Rename(oldPath, newPath)
			if err != nil {
				log.Fatal(err)
			}
		}
    }
}

/* Remove all file in dest. */
func RemoveFiles(dest Where) {
	destPath := GetWhere(dest)
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

/* Remove target file */
func RemoveFile(dest Where, filename string) {
	destPath := GetWhere(dest)
	files, err := os.ReadDir(destPath)
	if err != nil {
        log.Fatal(err)
    }
	for _, file := range files {
		if file.Name() == filename {
			filePath := filepath.Join(destPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				log.Fatal(err)
			}
		}
    }
}

func RemoveFileByPath(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

/* Find a file in directory, return filepath */
func FindFile(dir string, filename string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
        if file.Name() == filename {
            return filepath.Join(dir, file.Name())
        }
    }
	return ""
}

/* Read dir */
func ReadDir(dirname string) []fs.DirEntry {
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

/* If filename exist in dir */
func FileExists(filepath string) bool {
    _, err := os.Stat(filepath)
    if err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

/* If there some files in dir */
func DirHasFiles(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return len(files) > 0
}