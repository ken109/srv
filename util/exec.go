package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetOutput(name string, arg ...string) string {
	out, err := exec.Command(name, arg...).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.ReplaceAll(string(out), "\n", "")
}

func TryCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	} else {
		return nil
	}
}

func Cd(dir string) {
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dir
}

func Exists(name string) bool {
	f, err := os.Stat(name)
	return !(os.IsNotExist(err) || f.IsDir())
}

func Remove(name string) {
	if err := os.Remove(name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Mkdir(name string) {
	if err := os.Mkdir(name, 0755); err != nil {
		fmt.Println(err)
	}
}
