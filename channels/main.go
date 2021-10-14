package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	c := make(chan string)

	lou := []string{"http://google.com/", "http://yahoo.com/", "http://mohamadrezamotaghi.com/"}
	for _, l := range lou {
		go check(l, c)
		// fmt.Println(<-c)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(lou); i++ {
	// 	fmt.Println(<-c)
	// }

	for {
		go check(<-c, c)
	}

}

func check(l string, c chan string) {
	start := time.Now()
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "Is outOfReach")
		elapsed := time.Since(start).Seconds()
		log.Println(elapsed)
		c <- "DOWN"

		return
	}

	fmt.Println(l, "Is Ok!!")
	elapsed := time.Since(start).Seconds()
	log.Println(elapsed)

	c <- l + "   " + "UP"
}
