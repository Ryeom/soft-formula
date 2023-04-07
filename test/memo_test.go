package test

import (
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestMemo(t *testing.T) {
	cmd := exec.Command("cmd", "/c", "git add .")
	cmd.Dir = "D:\\dev\\windows-test"
	//fmt.Println("촤하하",cmd.Dir,"<-")
	cmd.Stdout = os.Stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "dir\n")
	}()
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
