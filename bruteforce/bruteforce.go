package bruteforce

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"
)

func isHashValid(hash string, target string) bool {
	fmt.Printf("Trying 0x%s\n", hash[:8])
	if len(hash) < len(target) {
		return false
	}

	for i, c := range target {
		if c == '*' { 
			continue
	 	}
		
		if c != rune(hash[i]) {
			return false
		}
	}
	return true
}

// functionName - *
func deeper(functionName string, target string, resultChannel chan<- string) {
	hash := sha3.NewLegacyKeccak256()

	cleanName := strings.Replace(functionName, "*", "", 1)
	
	hash.Write([]byte(cleanName))
	buf := hash.Sum(nil)

	if isHashValid(hex.EncodeToString(buf), target) {
		resultChannel <- cleanName
	} else {
		for _, letter := range alphabet {
			newName := strings.Replace(functionName, "*", letter + "*", 1)
			go deeper(newName, target, resultChannel)
		}
	}
}

func heartbeat() {
	fmt.Println("Still searching...")
  time.Sleep(2000)
}

func StartBruteforcing(functionName string, target string) {
	resultChannel := make(chan string)

	for _, letter := range alphabet {
		newName := strings.Replace(functionName, "*", letter + "*", 1)

		go deeper(newName, target, resultChannel)
	}

	for {
		select {
			case result := <-resultChannel:
				fmt.Println("Success! Result is ", result)
				return
			default: heartbeat()
		}
	}
}
