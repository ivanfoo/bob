package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"gopkg.in/urfave/cli.v1"
)

func Init(c *cli.Context) error {
	if c.Args().First() == "" {
		log.Fatal("A name in the format foo/bar is required")
	}

	remote := fmt.Sprintf("%s:%s", bobSrc, c.Args().First())
	err := addRemote(remote, c.String("remote"))

	if err != nil {
		fmt.Println("Can't setup git remote")
		return err
	}

	return nil
}

func addRemote(remote, name string) error {
	cmd := exec.Command("git", "remote", "add", name, remote)
	err := cmd.Run()

	return err
}
