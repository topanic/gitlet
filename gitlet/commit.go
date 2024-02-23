package gitlet

import (
	"encoding/json"
	"gitlet/config"
	"gitlet/utils"
	"log"
	"os"
	"time"
)


type Commit struct {
	Message string	`json:"message"`
	Parent *Commit	`json:"parents"`
	CurrDate time.Time	`json:"currDate"`
	HashId string	`json:"hashId"`
	BlobIds []string	`json:"blobIds"`
}

func NewCommit(message string) *Commit {
	blobs := GetStageBlog()
	blobIds := make([]string, 20)
	for _, b := range blobs {
		blobIds = append(blobIds, b.HashId)
	}
	return &Commit{
		Message: message,
		// TODO: parent
		CurrDate: time.Now(),
		HashId: utils.GenerateID(),
		BlobIds: blobIds,
	}
}

func (c *Commit) Persist() {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(config.COMMIT + "/" + c.HashId, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}