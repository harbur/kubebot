package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func execute(name string, arg ...string) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	cmd := exec.Command(name, arg...)

	stdin, err := cmd.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	stdin.Close()
	cmd.Wait()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
