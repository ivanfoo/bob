package cmd

import (
	"fmt"
	//"log"
	"os"
	"os/exec"

	"gopkg.in/urfave/cli.v1"
)

var (
	bobLibrary = os.Getenv("BOB_LIBRARY")
	bobSrc     = os.Getenv("BOB_SRC")
)

func Get(c *cli.Context) error {
	repo := fmt.Sprintf("%s/%s", bobSrc, c.Args().First())
	repoDir := fmt.Sprintf("%s/%s", bobLibrary, c.Args().First())

	if dirExists(repoDir) {
		fmt.Println("already cloned")
		return nil
	}

	err := cloneRepo(repo, repoDir)

	if err != nil {
		fmt.Println("error cloning repo")
	}

	return err
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	fmt.Println(path)
	if err == nil {
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		if os.IsExist(err) {
			fmt.Println("bbbbbb")
			return true
		}
	}

	return false
}

func cloneRepo(remote, path string) error {
	fmt.Println("cloning repository...")
	cmd := exec.Command("git", "clone", remote, path)
	err := cmd.Run()

	return err
}
