package instruction

import (
	"fmt"
	"gitlet/config"
	"gitlet/gitlet"
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