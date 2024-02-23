package instruction

import (
	"fmt"
	"gitlet/gitlet"
	"log"
	"os"
)

func Init_gitlet() {
	_, err := os.Stat(".gitlet")
	if os.IsNotExist(err) {
		// .getlet directory not exit
		os.Mkdir(".gitlet", 0755)
		os.MkdirAll(".gitlet/stage/addStage", 0755)
		os.MkdirAll(".gitlet/stage/rmStage", 0755)
		os.MkdirAll(".gitlet/objects/commits", 0755)
		os.MkdirAll(".gitlet/objects/blobs", 0755)
		os.MkdirAll(".gitlet/refs/heads", 0755)
		os.MkdirAll(".gitlet/refs/remotes", 0755)
		os.Create(".gitlet/HEAD")
		fmt.Println("Gitlet init success.")
	} else {
		// .gitlet directory exist
		fmt.Println("A Gitlet version-control system already exists in the current directory.")
	}
}

func Add(filenames ...string) {
	for _, filename := range filenames {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		blob := gitlet.NewBlob(filename, data)
		// store the blob
		blob.Persist()
	}	
	fmt.Println("Adding files succeed.")
}

func Commit(message string) {
	
}