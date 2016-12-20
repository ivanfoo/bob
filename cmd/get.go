package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/urfave/cli.v1"
)

var (
	bobSrc = os.Getenv("BOB_SRC")
)

func Get(c *cli.Context) error {
	remote := fmt.Sprintf("%s:%s", bobSrc, c.Args().First())
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, c.Args().First())

	if dirExists(path) {
		if c.Bool("update") {
			updateLocalRepo("forge", "master", path)
			return nil
		}

		fmt.Println("Repo already exists locally. To update it, run bob get -u " + c.Args().First())
		return nil
	}

	err := cloneRepo(remote, "forge", path)

	return err
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

func cloneRepo(remote, remoteName, path string) error {
	fmt.Print("Cloning repository...")

	cmd := exec.Command("git", "clone", "-o", remoteName, remote, path)
	err := cmd.Run()

	if err != nil {
		fmt.Println("error!")
		return err
	}

	fmt.Println("done")

	return nil
}

func updateLocalRepo(remote, branch, path string) error {
	fmt.Print("Updating repository...")

	err := os.Chdir(path)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "pull", remote, branch)
	err = cmd.Run()

	if err != nil {
		fmt.Println("error!")
		return err
	}

	fmt.Println("done")

	return nil
}
