package main

import (
	"fmt"
	"time"
)

func main() {
	seconds := 10

	loc := time.FixedZone("Asia/Tokyo", 9*60*60)

	fmt.Printf("Notfiy in %d seconds (Now: %v)\n", seconds, time.Now().In(loc))

	<-time.After(time.Duration(seconds) * time.Second)
	fmt.Printf("Finish! (Now: %v)\n", time.Now().In(loc))
}
