package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Find in URL string sequence of non-slashes ending with RPM, DEB or tar.<arch>
	// URL can continue after extension (I have not seen such with package repositiories
	// but URI syntax allows to append queries or fragments after path)
	re := regexp.MustCompile(`[^\/]+\.(rpm|u?deb|tar\.(gz|bz2|xz))`)

	for true {
		line, _ := reader.ReadString('\n')
		parts := strings.Split(line, " ")
		url := parts[0]
		filename := re.FindString(url)

		if len(filename) > 0 {
			// Success. A new storage ID is presented for this URL.
			fmt.Printf("OK store-id=%v\n", filename)
		} else {
			// Success. No change for this URL.
			fmt.Printf("ERR\n")
		}
	}
}
