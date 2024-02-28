package main

import (
	"fmt"
	"gitlet/instruction"
	"gitlet/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("You need input at least an argument.")
		return
	}

	switch args[0] {
	case "init":
		instruction.Init_gitlet()
	case "add":
		if utils.GetArgsNum(args) == 2 {
			instruction.Add(args[1])
		} else {
			fmt.Println("add: You need input at least a file.")
		}
	case "commit":
		if utils.GetArgsNum(args) >= 2 {
			instruction.Commit(args[1:]...)
		} else {
			fmt.Println("commit: Get wrong argument num.")
		}
	case "rm":
		if utils.GetArgsNum(args) == 2 {
			instruction.Rm(args[1])
		} else {
			fmt.Println("rm: Get wrong argument num.")
		}
	case "log":
		if utils.GetArgsNum(args) == 1 {
			instruction.Log()
		} else {
			fmt.Println("log: Get wrong argument num.")
		}
	case "global-log":
		if utils.GetArgsNum(args) == 1 {
			instruction.GlobalLog()
		} else {
			fmt.Println("log: Get wrong argument num.")
		}
	case "find":
		if utils.GetArgsNum(args) >= 2 {
			instruction.Find(args[1:]...)
		} else {
			fmt.Println("find: Get wrong argument num.")
		}
	case "status":
		if utils.GetArgsNum(args) == 1 {
			instruction.Status()
		} else {
			fmt.Println("status: Get wrong argument num.")
		}
	case "checkout":
		if utils.GetArgsNum(args) >= 2 {
			instruction.Checkout(args[1:]...)
		} else {
			fmt.Println("checkout: Get wrong argument num.")
		}
	case "branch":
		if utils.GetArgsNum(args) == 2 {
			instruction.Branch(args[1])
		} else {
			fmt.Println("branch: Get wrong argument num.")
		}



	default:
		fmt.Println("Please input a valid instruction.")
		return
	}
}