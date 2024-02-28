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

/* Change HEAD point to */
func MoveHEAD(branchName string) {
	branchPath := config.BRANCHES + "/" + branchName
	utils.WriteFile(config.HEAD, branchPath)
}

/* Move Branch point to. */
func MoveBranchPoint(commitId string) {
	data := utils.ReadFile(config.HEAD)
	refs_branch := string(data)
	// blobId, err := os.ReadFile(refs_branch)
	utils.WriteFile(refs_branch, commitId)
}

func BranchExist(branchName string) bool {
	return utils.FileExists(config.BRANCHES + "/" + branchName)
}