package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check(e error) {
	// Reading files requires checking most calls for errors.
	// This helper will streamline our error checks below.
	if e != nil {
		panic(e)
	}
}

func userInput(stubCfg *stubConfig) (string, string) {
	var srcFilePath string
	fmt.Print("[?] Path of file to be encrypted: ")
	fmt.Scanln(&srcFilePath)

	var dstFilePath string
	fmt.Print("[?] Path of output (encrypted) file: ")
	fmt.Scanln(&dstFilePath)

	fmt.Print("[?] Name of the target process: ")
	fmt.Scanln(&stubCfg.ProcName)

	fmt.Print("[?] Encryption key (32 bits - random if empty): ")
	fmt.Scanln(&stubCfg.EncryptionIV)
	fmt.Print("[?] Encryption IV (16 bits - random if empty): ")
	fmt.Scanln(&stubCfg.EncryptionIV)
	if stubCfg.EncryptionKey == "" {
		stubCfg.EncryptionKey = randKey(modeKey)
		stubCfg.EncryptionIV = randKey(modeIV)
	}
	fmt.Println()
	fmt.Printf("[!] Random encryption key (used in stub): %s\n", stubCfg.EncryptionKey)
	fmt.Printf("[!] Random encryption IV (used in stub): %s\n", stubCfg.EncryptionIV)
	return srcFilePath, dstFilePath
}

func randKey(mode int) string {
	var keySize int

	if mode == modeIV {
		keySize = 16
	} else if mode == modeKey {
		keySize = 32
	}

	key := make([]byte, keySize)
	for i := range key {
		key[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return string(key)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
