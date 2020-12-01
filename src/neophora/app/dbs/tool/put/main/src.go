package main

import (
	"encoding/hex"
	"log"
	"neophora/cli"
	"os"
	"strings"
)

func main() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	addresses := strings.Split(address, " ")
	client := &cli.T{
		Addresses: addresses,
		TryTimes:  3,
	}

	key := []byte(os.Args[1])
	value := []byte(os.Args[2])
	var ret bool
	var err error

	switch os.ExpandEnv("${DEC}") {
	case "hex":
		value, err = hex.DecodeString(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
	}

	if err := client.Call("DB.Put", struct {
		Key   []byte
		Value []byte
	}{
		Key:   key,
		Value: value,
	}, &ret); err != nil {
		log.Fatalln(err)
	}

	log.Println(ret)
}