package main

import (
	"log"
	"net/rpc"
	"os"
)

func main() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	var ret bool
	err = client.Call("DB.Put", struct {
		Key   []byte
		Value []byte
	}{
		Key:   []byte(os.Args[1]),
		Value: []byte(os.Args[2]),
	}, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ret)
}
