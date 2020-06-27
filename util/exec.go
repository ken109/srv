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
	_, err := os.Stat(name)
	return err == nil
}

func Remove(name string) {
	if err := os.Remove(name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
