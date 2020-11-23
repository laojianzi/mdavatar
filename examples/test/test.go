package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/laojianzi/mdavatar"
)

func main() {
	initials := "LR"
	avatar, err := mdavatar.New(initials).Build()
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("out-%d.png", time.Now().Unix())
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(file, avatar); err != nil {
		panic(err)
	}
}
