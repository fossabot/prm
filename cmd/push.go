package cmd

import (
	"fmt"
	"os"

	"github.com/ldez/prm/config"
	"github.com/ldez/prm/local"
	"github.com/ldez/prm/types"
)

// Push push to the PR branch.
func Push(options *types.PushOptions) error {

	// get configuration
	confs, err := config.ReadFile()
	if err != nil {
		return err
	}

	repoDir, err := os.Getwd()
	if err != nil {
		return err
	}

	con, err := config.Find(confs, repoDir)
	if err != nil {
		return err
	}

	number, err := local.GetCurrentPRNumber(options.Number)
	if err != nil {
		return err
	}

	pr, err := con.FindPullRequests(number)
	if err != nil {
		return err
	}

	fmt.Println("push", pr)

	err = pr.Push(options.Force)
	if err != nil {
		return err
	}

	return nil
}
