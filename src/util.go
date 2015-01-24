package hello

import (
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func extendMap(a map[string]Gif, b map[string]Gif) {
	for k, v := range b {
		a[k] = v
	}
}
