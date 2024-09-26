package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

const (
	stubDir      = "stub"
	allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#$%0123456789"
	modeIV       = iota
	modeKey
)

type stubConfig struct {
	ProcName      string
	EncryptionKey string
	EncryptionIV  string
}

func createStub(stubCfg *stubConfig) []byte {
	f, err := os.Create("stub/vars.go")
	check(err)
	defer f.Close()

	tmpl, err := template.New("").Parse(`// Code generated automatically; DO NOT EDIT.
// Generated using data from user input
package main

var (
	key      = "{{.EncryptionKey}}"
	iv       = "{{.EncryptionIV}}"
	procName = "{{.ProcName}}"
)
`)
	check(err)
	tmpl.Execute(f, stubCfg)
	os.Chdir(stubDir)
	cmdOut, err := exec.Command("go", "build", "-ldflags", `-w -s`, ".").CombinedOutput()
	if len(cmdOut) > 0 {
		fmt.Println("Output: ", string(cmdOut))
	}
	check(err)
	stubBytes, err := ioutil.ReadFile("stub")
	check(err)
	os.Chdir("..")

	return stubBytes
}

func main() {
	stubCfg := &stubConfig{}
	srcFilePath, dstFilePath := userInput(stubCfg)

	srcBytes, err := ioutil.ReadFile(srcFilePath)
	check(err)
	encryptedBytes := aesEnc(srcBytes, stubCfg.EncryptionKey, stubCfg.EncryptionIV)

	fmt.Println("[!] Generating stub...")
	stubBytes := createStub(stubCfg)

	fmt.Println("[!] Creating final executable...")
	file, err := os.Create(dstFilePath)
	check(err)
	w := bufio.NewWriter(file)

	w.Write(stubBytes)
	w.Write([]byte(stubCfg.EncryptionKey))
	w.Write([]byte(stubCfg.EncryptionIV))
	w.Write(encryptedBytes)
	w.Flush()
	fmt.Println("[!] All done!")
}
