package storage

import (
	"bufio"
	"fmt"
	"github.com/sadek-39/key-value-database/types"
	"os"
	"strings"
	"sync"
)

var StoreFile = "store.txt"
var FileMutex = &sync.Mutex{}
var Store = make(types.StoreType)

func SaveDataToFile(store types.StoreType) {
	FileMutex.Lock()
	defer FileMutex.Unlock()

	file, err := os.OpenFile(StoreFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error when opening file : ", err)
		_, cerr := os.Create(StoreFile)
		if cerr != nil {
			fmt.Println("Error from creating file:  ", cerr)
		}
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for key, value := range store {
		_, err := fmt.Fprintf(writer, "%s=%s\n", key, value)
		if err != nil {
			return
		}
	}

	werr := writer.Flush()
	if werr != nil {
		return
	}
}

func LoadDataFromFile() {
	FileMutex.Lock()
	defer FileMutex.Unlock()

	file, err := os.Open(StoreFile)
	if err != nil {
		fmt.Println("Error when opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")

		if len(parts) != 2 {
			fmt.Println("Skipping malformed line:", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		Store[key] = value
	}
}

func Get(key string) (string, bool) {
	FileMutex.Lock()
	defer FileMutex.Unlock()

	value, found := Store[key]
	return value, found
}
