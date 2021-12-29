package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

//go:embed script.sh.b64
var scriptB64 string

func Exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func main() {
	//var arguments = strings.Join(os.Args[1:], " ")
	var shell = ""
	if Exists("/bin/sh") {
		shell = "/bin/sh"

	} else if Exists("/bin/sh") {
		shell = "/bin/sh"

	} else if Exists("/bin/zsh") {
		shell = "/bin/zsh"

	} else if Exists("/bin/ksh") {
		shell = "/bin/ksh"

	} else if Exists("/bin/csh") {
		shell = "/bin/csh"

	} else {
		fmt.Println("No shell was found, I cannot execute the script.")
		return
	}

	//if len(arguments) > 0 {
	//	shell = shell + " -s -- " + arguments
	//	fmt.Println("Executing with arguments: " + arguments)
	//}

	c1 := exec.Command("base64", "-d")
	c2 := exec.Command(shell)

	c1.Stdin = strings.NewReader(scriptB64)
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
}
