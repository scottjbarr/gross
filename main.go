package main

import (
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Printf("Usage : %v url\n", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	url := &os.Args[1]

	client := new(Client)

	rss, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	log.Printf("%v\n", rss.Channel.Title)

	for _, item := range rss.Channel.Items {
		log.Printf("%v (%v)\n", item.Title, item.Link)
	}
}
