package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	var typeSHA string
	flag.StringVar(&typeSHA, "type-SHA", "SHA256", "type SHA")
	flag.Parse()

	userString := readString()
	c1 := sha256.Sum256([]byte("x"))
	countBit(c1)
	c2 := sha256.Sum256([]byte("X"))
	countBit(c2)

	switch typeSHA {
	case "SHA384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(userString)))
	case "SHA512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(userString)))
	default:
		sha := sha512.Sum512_256([]byte(userString))
		fmt.Printf("%x\n", sha)
		countBit(sha)
	}
}

func countBit(sha [32]byte) {
	var countBit int
	for _, v := range sha {
		countBit += bits.OnesCount64(uint64(v))
	}

	fmt.Println(countBit)
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
