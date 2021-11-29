package dictionary

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error while closing file : %v\n", err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PickWord() string {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}
