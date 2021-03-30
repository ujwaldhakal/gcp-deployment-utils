package git

import (
	"os"
	"os/exec"
	"strings"
)

func GetCommitHash() string {
	commitHash, err := exec.Command("git", "rev-parse", "HEAD").Output()

	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSuffix(string(commitHash), "\n")
}