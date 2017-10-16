package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type song struct {
	Name     string
	Duration string
}

type songs struct {
	Song []song
}

func main() {
	var favorite songs

	if _, err := toml.DecodeFile("fixture/songs.toml", &favorite); err != nil {
		log.Fatal(err)
	}

	for _, s := range favorite.Song {
		fmt.Printf("%s (%s)\n", s.Name, s.Duration)

	}

}
