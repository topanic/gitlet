package instruction

import (
	"fmt"
	"gitlet/config"
	"gitlet/gitlet"
	"gitlet/utils"
	"log"
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
		// write commitADD
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

	// 1. copy Hashmap from HEAD
	commitId := gitlet.GetHEAD()
	oldCommit := gitlet.GetCommitById(commitId)
	blobs := gitlet.GetStageBlob(utils.ADDSTAGE)
	// Deep and shallow copies may not matter here, 
	// because persistence will read information from the local file,
	// and what is read out is still what was originally stored in the file.
	blobIds := oldCommit.BlobIds

	// 2. add blob into Hashmap from "addStage"
	for _, b := range blobs {
		// blobIds = append(blobIds, b.HashId)
		if _, ok := blobIds[b.FilePath]; !ok {
			blobIds[b.FilePath] = b.HashId
		} else {
			log.Fatalln("store same filepath, something get wrong")
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
		delete(commit.BlobIds, rmblob.FilePath)
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
		} else {
			// 3. if file doesn't exit in worktree but is traced by commit, 
			//    move it from "objects/blobs" to "rmStage".
			utils.MoveFile(utils.BLOB, utils.RMSTAGE, blobId)
			fmt.Println("rm file from commit.")
		}
	}
	
}