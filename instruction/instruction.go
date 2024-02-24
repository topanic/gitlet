package instruction

import (
	"fmt"
	"gitlet/config"
	"gitlet/gitlet"
	"gitlet/utils"
	"os"
)

func Init_gitlet() {
	_, err := os.Stat(".gitlet")
	if os.IsNotExist(err) {
		// .getlet directory not exit
		os.Mkdir(".gitlet", 0755)
		os.MkdirAll(config.ADDSTAGE, 0755)
		os.MkdirAll(config.RMSTAGE, 0755)
		os.MkdirAll(config.COMMIT, 0755)
		os.MkdirAll(config.BLOB, 0755)
		os.MkdirAll(config.BRANCHES, 0755)
		os.MkdirAll(config.REMOTES, 0755)
		os.Create(config.HEAD)
		os.Create(config.BRANCHES + "/master")
		// write commitId into "refs/heads/master"
		commit := gitlet.NewInitCommit()
		utils.WriteFile(config.BRANCHES + "/master", commit.HashId)
		// write "refs/heads/master" into HEAD
		utils.WriteFile(config.HEAD, config.BRANCHES + "/master")
		// write commit
		commit.Persist()
		fmt.Println("Gitlet init success.")
	} else {
		// .gitlet directory exist
		fmt.Println("A Gitlet version-control system already exists in the current directory.")
	}
}

func Add(filenames ...string) {
	for _, filename := range filenames {
		data := utils.ReadFile(filename)
		blob := gitlet.NewBlob(filename, data)
		// store the blob
		blob.Persist()
	}	
	fmt.Println("Adding files succeed.")
}

func Commit(message string) {
	commit := gitlet.NewCommit(message)
	// store this commit
	commit.Persist()
	// store blobs which contains in the commit blobs into "objects/blobs"
	utils.MoveFiles(utils.ADDSTAGE, utils.BLOB)
	// remove "addStage" file
	utils.RemoveFiles(utils.ADDSTAGE)
	// move HEAD
	gitlet.MoveHEAD(commit.HashId)
	fmt.Println("Commit succeed.")
}