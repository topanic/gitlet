package gitlet

import (
	"gitlet/config"
	"log"
	"os"
)

/* Get commitId which HEAD point to. */
func GetHEAD() string {
	data, err := os.ReadFile(config.HEAD)
	if err != nil {
		log.Fatal(err)
	}
	refs_branch := string(data)
	blobId, err := os.ReadFile(refs_branch)
	if err != nil {
		log.Fatal(err)
	}
	return string(blobId)
}

/* Move HEAD point to. */
func MoveHEAD(commitId string) {
	data, err := os.ReadFile(config.HEAD)
	if err != nil {
		log.Fatal(err)
	}
	refs_branch := string(data)
	// blobId, err := os.ReadFile(refs_branch)
	err = os.WriteFile(refs_branch, []byte(commitId), 0644)
	if err != nil {
		log.Fatal(err)
	}
}