package gitlet

import (
	"encoding/json"
	"gitlet/config"
	"gitlet/utils"
	"log"
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

func (b *Blob) Persist() {
	data, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteFileBytes(config.ADDSTAGE + "/" + b.HashId, data)
}


func GetStageBlob(w utils.Where) []*Blob {
	// TODO: Need refactor.
	dirname := utils.GetWhere(w)
	files := utils.ReadDir(dirname)
	blobs := make([]*Blob, 0, 10)
	for _, file := range files {
		filepath := dirname + "/" + file.Name()
		data := utils.ReadFile(filepath)
		b := &Blob{}
		err := json.Unmarshal(data, b)
		if err != nil {
			log.Fatal(err)
		}
		blobs = append(blobs, b)
	}
	return blobs
}

func GetBlobByFilename(filename string, w utils.Where) *Blob {
	dirname := utils.GetWhere(w)
	files := utils.ReadDir(dirname)
	for _, file := range files {
		if file.Name() == filename {
			data := utils.ReadFile(filename)
			b := &Blob{}
			err := json.Unmarshal(data, b)
			if err != nil {
				log.Fatal(err)
			}
			return b
		}
	}
	// can't find blob named "filename".
	return nil
}