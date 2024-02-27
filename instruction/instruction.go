package instruction

import (
	"bytes"
	"fmt"
	"gitlet/config"
	"gitlet/gitlet"
	"gitlet/utils"
	"log"
	"os"
	"strings"
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
		// write commitADD
		commit.Persist()
		fmt.Println("Gitlet init success.")
	} else {
		// .gitlet directory exist
		fmt.Println("A Gitlet version-control system already exists in the current directory.")
	}
}

func Add(filename string) {
	data := utils.ReadFile(filename)
	blob := gitlet.NewBlob(filename, data)
	addBlobs := gitlet.GetStageBlob(utils.ADDSTAGE)
	for _, b := range addBlobs {
		if b.FilePath == blob.FilePath {
			if bytes.Equal(b.Contents, blob.Contents) {
				// same file
				fmt.Println("add: file is already added.")
				return
			} else {
				// changed file
				utils.RemoveFile(utils.ADDSTAGE, b.HashId)
				break
			}
		}
	}
	// store the blob
	blob.Persist()
	
	fmt.Println("Adding files succeed.")
}

func Commit(messages ...string) {
	message := strings.Join(messages, " ")
	// if some files exist in stage
	if !utils.DirHasFiles(config.ADDSTAGE) && !utils.DirHasFiles(config.RMSTAGE) {
		fmt.Println("Commit: Nothing to do.")
		return
	}
	commit := gitlet.NewCommit(message)
	// 1. copy Hashmap from HEAD
	commitId := gitlet.GetHEAD()
	oldCommit := gitlet.GetCommitById(commitId)
	// Deep and shallow copies may not matter here, 
	// because persistence will read information from the local file,
	// and what is read out is still what was originally stored in the file.
	blobIds := oldCommit.BlobIds

	// 2. add blob into Hashmap from "addStage"
	addBlobs := gitlet.GetStageBlob(utils.ADDSTAGE)
	for _, addBlob := range addBlobs {
		if blobID, ok := blobIds[addBlob.FilePath]; !ok {
			blobIds[addBlob.FilePath] = addBlob.HashId
		} else {
			// file exist in commit, compare the content,
			// if content is same, remove it from "addStage"
			// if content is not same, rewrite Hashmap.
			objBlob := gitlet.GetBlobById(blobID, utils.BLOB)
			if objBlob != nil {
				if bytes.Equal(addBlob.Contents, objBlob.Contents) {
					// content same
					utils.RemoveFile(utils.ADDSTAGE, addBlob.HashId)
				} else {
					// content not same
					blobIds[addBlob.FilePath] = addBlob.HashId
				}
			} else {
				log.Fatalln("you should't arrive here, some thing get wrong.")
			}
		}
	}

	// 3. move "addStage" to "objects/blobs"
	// store blobs which contains in the commit blobs into "objects/blobs"
	utils.MoveFiles(utils.ADDSTAGE, utils.BLOB)

	// 4. remove "addStage" file
	utils.RemoveFiles(utils.ADDSTAGE)
	
	// 5. remove "rmStage" from Hashmap
	rmBlobs := gitlet.GetStageBlob(utils.RMSTAGE) 
	for _, rmblob := range rmBlobs {
		delete(blobIds, rmblob.FilePath)
	}

	// 6. remove "rmStage" file
	utils.RemoveFiles(utils.RMSTAGE)

	// 7. store this commit and move HEAD
	// store this commit
	commit.BlobIds = blobIds
	commit.Persist()
	// move HEAD
	gitlet.MoveHEAD(commit.HashId)
	fmt.Println("Commit succeed.")
}

func Rm(filename string) {
	// 1. if file exist in "addStage" directory, remove it directly.
	blob := gitlet.GetBlobByFilename(filename, utils.ADDSTAGE)
	if blob != nil {
		err := os.Remove(config.ADDSTAGE + "/" + blob.HashId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("rm file from addStage.")
		return
	}
	
	fileExist := utils.FileExists(filename)
	commitId := gitlet.GetHEAD()
	commit := gitlet.GetCommitById(commitId)
	if blobId, ok := commit.BlobIds[filename]; ok {
		if fileExist {
			// 2. if file exits in worktree and is traced by commit, 
			//    move it in "rmStage" and remove it in worktree.
			utils.MoveFile(utils.BLOB, utils.RMSTAGE, blobId)
			err := os.Remove(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rm file from commit and worktree.")
			return
		} else {
			// 3. if file doesn't exit in worktree but is traced by commit, 
			//    move it from "objects/blobs" to "rmStage".
			utils.MoveFile(utils.BLOB, utils.RMSTAGE, blobId)
			fmt.Println("rm file from commit.")
			return
		}
	}
	fmt.Println("rm: Nothing to do.")
}

func Log() {
	commitIdLast := gitlet.GetHEAD()
	logHelper(commitIdLast)
}

func logHelper(commitId string) {
	commit := gitlet.GetCommitById(commitId)
	fmt.Printf("* %s %s\n", utils.Colorize(commit.HashId[0:7], utils.FgMagenta),
		utils.Colorize(commit.Message, utils.FgCyan))
	if commit.Parent != nil {
		logHelper(commit.Parent[0])
	}
}

func GlobalLog() {
	dirs := utils.ReadDir(config.COMMIT)
	for _, item := range dirs {
		fmt.Printf("* ")
		fmt.Println(utils.Colorize(item.Name()[:7], utils.FgMagenta))
	}
}

func Find(commitMessage ...string) {
	commits := gitlet.GetAllCommits()
	for _, commit := range commits {
		if commit.Message == strings.Join(commitMessage, " ") {
			fmt.Println(utils.Colorize(commit.HashId[:7], utils.FgMagenta))
		}
	}
}

func Status() {
	fmt.Printf("=== Branches ===\n")
	HEADBranch := gitlet.GetHEADBranch()
	branches := utils.ReadDir(config.BRANCHES)
	fmt.Printf("*%s\n", HEADBranch)
	for _, branch := range branches {
		if name := branch.Name(); name != HEADBranch {
			fmt.Printf("*%s\n", name)
		}
	}

	fmt.Printf("=== Staged Files ===\n")
	stages := gitlet.GetStageBlob(utils.ADDSTAGE)
	for _, stage := range stages {
		fmt.Println(stage.Filename)
	}

	fmt.Printf("=== Removed Files ===\n")
	stages = gitlet.GetStageBlob(utils.RMSTAGE)
	for _, stage := range stages {
		fmt.Println(stage.Filename)
	}

	// fmt.Printf("=== Modifications Not Staged For Commit ===\n")
	// TODO

	// fmt.Printf("=== Untracked Files ===\n")
	// TODO
}
