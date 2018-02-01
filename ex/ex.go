package main

import (
	"fmt"
	"github.com/as/clip"
	"log"
)

func no(s string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", s, err)
	}
}

func main() {
	c, err := clip.New()
	no("clip.New", err)
	defer c.Close()

	n, err := c.Write([]byte("T\x00e\x00s\x00t\x00\x00\x00"))
	no("c.Write", err)

	buf := make([]byte, 1024)
	p := make([]byte, 1024)
	for i := 0; i < 1000; i++ {
		go func() {
			n, err = c.Read(buf)
			no("c.Read 1", err)
			buf = buf[:n]

			n, err = c.Read(p)
			no("c.Read 2", err)
		}()
	}
	fmt.Println("clipboard:", string(p))
}
