package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/laojianzi/mdavatar"
	"github.com/laojianzi/mdavatar/style"
)

func main() {
	avatar, err := mdavatar.New("MDAvatar").Builds(style.NewCircle)
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("mdavatar-circle-%d.png", time.Now().Unix())
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(file, avatar); err != nil {
		log.Fatal(err)
	}
}
