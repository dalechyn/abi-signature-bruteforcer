package bruteforce

import (
	"encoding/hex"
	"fmt"
	"strings"

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

func StartBruteforcing(functionName string, target string) {
	for combination := range GenerateCombinations(8) {
		hash := sha3.NewLegacyKeccak256()

		cleanName := strings.Replace(functionName, "*", combination, 1)
		
		hash.Write([]byte(cleanName))
		buf := hash.Sum(nil)

		if isHashValid(hex.EncodeToString(buf), target) {
			fmt.Println("Success! Result is ", cleanName)
			return
		} 
	}
}

// 0
// 00
// 000

// 1
// ..
// X
//.. -> 
//   

// |x| <-32 = 32
// |x|x| <-32 = 32^2
// |x|x|x| <-32 = 32^3