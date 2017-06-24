package main

import (
	"fmt"
	"log"
	"os"

	"github.com/containous/flaeg"
	"github.com/ldez/prm/cmd"
	"github.com/ldez/prm/types"
)

func main() {

	emptyConfig := &types.NoOption{}
	rootCmd := &flaeg.Command{
		Name:                  "prm",
		Description:           "PRM - The Pull Request Manager.",
		Config:                emptyConfig,
		DefaultPointersConfig: &types.NoOption{},
		Run: func() error {
			return nil
		},
	}

	flag := flaeg.New(rootCmd, os.Args[1:])

	// Checkout

	checkoutOptions := &types.CheckoutOptions{}

	checkoutCmd := &flaeg.Command{
		Name:                  "c",
		Description:           "Checkout a PR (create a local branch and add remote).",
		Config:                checkoutOptions,
		DefaultPointersConfig: &types.CheckoutOptions{},
	}
	checkoutCmd.Run = func() error {
		err := requirePRNumber(checkoutOptions.Number, checkoutCmd.Name)
		if err != nil {
			log.Fatalln(err)
		}
		err = cmd.Checkout(checkoutOptions)
		if err != nil {
			log.Println(err)
		}
		return nil
	}

	flag.AddCommand(checkoutCmd)

	// Remove

	removeOptions := &types.RemoveOptions{}

	removeCmd := &flaeg.Command{
		Name:                  "rm",
		Description:           "Remove one or more PRs from the current local repository.",
		Config:                removeOptions,
		DefaultPointersConfig: &types.RemoveOptions{},
	}
	removeCmd.Run = func() error {
		err := requirePRNumber(removeOptions.Number, removeCmd.Name)
		if !removeOptions.All && err != nil {
			log.Fatalln(err)
		}
		err = cmd.Remove(removeOptions)
		if err != nil {
			log.Println(err)
		}
		return nil
	}

	flag.AddCommand(removeCmd)

	// Push Force

	pushForceOptions := &types.PushOptions{Force: true}

	pushForceCmd := &flaeg.Command{
		Name:                  "pf",
		Description:           "Push force to the PR branch.",
		Config:                pushForceOptions,
		DefaultPointersConfig: &types.PushOptions{},
	}
	pushForceCmd.Run = func() error {
		err := cmd.Push(pushForceOptions)
		if err != nil {
			log.Println(err)
		}
		return nil
	}

	flag.AddCommand(pushForceCmd)

	// Push

	pushOptions := &types.PushOptions{}

	pushCmd := &flaeg.Command{
		Name:                  "push",
		Description:           "Push to the PR branch.",
		Config:                pushOptions,
		DefaultPointersConfig: &types.PushOptions{},
	}
	pushCmd.Run = func() error {
		err := cmd.Push(pushOptions)
		if err != nil {
			log.Println(err)
		}
		return nil
	}

	flag.AddCommand(pushCmd)

	// List

	listOptions := &types.ListOptions{}

	listCmd := &flaeg.Command{
		Name:                  "list",
		Description:           "Display all current PRs.",
		Config:                listOptions,
		DefaultPointersConfig: &types.ListOptions{},
		Run: func() error {
			err := cmd.List(listOptions)
			if err != nil {
				log.Println(err)
			}
			return nil
		},
	}

	flag.AddCommand(listCmd)

	// Run command
	flag.Run()
}

func requirePRNumber(number int, action string) error {
	if number <= 0 {
		return fmt.Errorf("You must provide a PR number. ex: 'prm %s -n 1235'", action)
	}
	return nil
}
