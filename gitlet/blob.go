package gitlet

import (
	"encoding/json"

	"gitlet/config"
	"gitlet/utils"
	"log"
	"os"
	"path/filepath"
)


type Blob struct {
	Filename string `json:"filename"`
	FilePath string `json:"filepath"`
	Contents []byte `json:"contents"`
	HashId   string `json:"hashId"`
}


func NewBlob(filePath string, contents []byte) *Blob {
	return &Blob{
		Filename: filepath.Base(filePath),
		FilePath: filePath,
		Contents: contents,
		HashId: utils.GenerateID(),
	}
}

func GetStageBlog() []*Blob {
	dirname := config.ADDSTAGE
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	blobs := make([]*Blob, 0, 10)
	for _, file := range files {
		filepath := config.ADDSTAGE + "/" + file.Name()
		data := utils.ReadFile(filepath)
		b := &Blob{}
		err = json.Unmarshal(data, b)
		if err != nil {
			log.Fatal(err)
		}
		blobs = append(blobs, b)
	}
	// TODO: need rm rmstage file
	return blobs
}



func (b *Blob) Persist() {
	data, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteFileBytes(config.ADDSTAGE + "/" + b.HashId, data)
}

