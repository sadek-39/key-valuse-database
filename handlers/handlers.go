package handlers

import (
	"fmt"
	"github.com/sadek-39/key-value-database/storage"
	"github.com/sadek-39/key-value-database/types"
	"strings"
)

var store = make(types.StoreType)

var Handlers = map[string]func(string){
	"ping": func(args string) {
		ping()
	},
	"set": func(args string) {
		set(args)
	},
	"get": func(args string) {
		get(args)
	},
}

func set(args string) {
	text := strings.TrimSpace(args)
	parts := strings.SplitN(text, " ", 2)

	if len(parts) != 2 {
		fmt.Println("Invalid arguments for set command")
	}

	key := parts[0]
	value := parts[1]

	store[key] = value

	storage.SaveDataToFile(store)

	fmt.Println("store: ", store)
}

func ping() {
	fmt.Println("pong")
}

func get(args string) {
	text := strings.TrimSpace(args)

	fmt.Println("Get for key : ", text)

	if text == "" {
		fmt.Println("Invalid arguments for get command")
	}

	if value, found := storage.Get(text); found {
		fmt.Println("value is: ", value)
	} else {
		fmt.Println("Key not found")
	}
}
