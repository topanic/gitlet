package gitlet

import (
	"gitlet/config"
	"gitlet/utils"
	"path/filepath"
)

/* Get commitId which HEAD point to. */
func GetHEAD() string {
	data := utils.ReadFile(config.HEAD)
	refs_branch := string(data)
	blobId := utils.ReadFile(refs_branch)
	return string(blobId)
}

/* Get HEAD Branch */
func GetHEADBranch() string {
	data := utils.ReadFile(config.HEAD)
	refs_branch := string(data)
	return filepath.Base(refs_branch)
}

/* Move HEAD point to. */
func MoveHEAD(commitId string) {
	data := utils.ReadFile(config.HEAD)
	refs_branch := string(data)
	// blobId, err := os.ReadFile(refs_branch)
	utils.WriteFile(refs_branch, commitId)
}