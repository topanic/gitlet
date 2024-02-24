package gitlet

import (
	"encoding/json"
	"gitlet/config"
	"gitlet/utils"
	"log"
	"time"
)


type Commit struct {
	Message string	`json:"message"`
	Parent []string	`json:"parents"`
	CurrDate time.Time	`json:"currDate"`
	HashId string	`json:"hashId"`
	BlobIds map[string]string	`json:"blobIds"`
}

func NewCommit(message string) *Commit {
	blobs := GetStageBlog()
	blobIds := make(map[string]string, 10)
	for _, b := range blobs {
		// blobIds = append(blobIds, b.HashId)
		if _, ok := blobIds[b.FilePath]; !ok {
			blobIds[b.FilePath] = b.HashId
		} else {
			log.Fatalln("store same filepath, something get wrong")
		}
	}
	return &Commit{
		Message: message,
		Parent: []string{GetHEAD()},
		CurrDate: time.Now(),
		HashId: utils.GenerateID(),
		BlobIds: blobIds,
	}
}

/* Init commit */
func NewInitCommit() *Commit {
	return &Commit{
		Message: "Init Commit",
		Parent: nil,
		CurrDate: time.Now(),
		HashId: utils.GenerateID(),
		BlobIds: nil,
	}
}

func (c *Commit) Persist() {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteFileBytes(config.COMMIT + "/" + c.HashId, data)
}