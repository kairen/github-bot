package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	head   = "git"
	path   = "/tmp/github-bot"
	remote = "kairen"
	repos  = "https://github.com/kairen/github-bot.git"
)

func TestGitClone(t *testing.T) {
	GitClone(path, repos)
	_, err := os.Stat(path + "github-bot")
	if !os.IsNotExist(err) {
		log.Fatal("Git clone failed.")
	}
}

func TestGitAddRemote(t *testing.T) {
	GitAddRemote(path, remote, repos)
	args := strings.Split(fmt.Sprintf("-C %s remote show", path), " ")
	out, _ := exec.Command(head, args...).Output()
	if !strings.Contains(string(out), remote) {
		log.Fatalf("Git add remote failed: %s", string(out))
	}
}
