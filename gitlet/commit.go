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
	return &Commit{
		Message: message,
		Parent: []string{GetHEAD()},
		CurrDate: time.Now(),
		HashId: utils.GenerateID(),
		BlobIds: nil,
	}
}

/* Init commit */
func NewInitCommit() *Commit {
	return &Commit{
		Message: "Init Commit",
		Parent: nil,
		CurrDate: time.Now(),
		HashId: utils.GenerateID(),
		BlobIds: make(map[string]string),
	}
}

func (c *Commit) Persist() {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteFileBytes(config.COMMIT + "/" + c.HashId, data)
}


/* Get commit entity by commitId */
func GetCommitById(id string) *Commit {
	commit := &Commit{}
	filepath := utils.FindFile(config.COMMIT, id)
	if filepath == "" {
		log.Fatalln("can't find file, something get wrong.")
	}
	data := utils.ReadFile(filepath)
	err := json.Unmarshal(data, commit)
	if err != nil {
		log.Fatal(err)
	}
	return commit
}

/* Get all commits */
func GetAllCommits() []*Commit {
	commits := make([]*Commit, 0)
	dirs := utils.ReadDir(config.COMMIT)
	for _, item := range dirs {
		filePath := config.COMMIT + "/" + item.Name()
		data := utils.ReadFile(filePath)
		commit := &Commit{}
		err := json.Unmarshal(data, commit)
		if err != nil {
			log.Fatal(err)
		}
		commits = append(commits, commit)
	} 
	return commits
}
