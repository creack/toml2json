package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	m := map[string]interface{}{}

	_, err := toml.DecodeReader(os.Stdin, &m)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.NewEncoder(os.Stdout).Encode(m)
}
