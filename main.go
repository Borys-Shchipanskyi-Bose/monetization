package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BoseCorp/monetization/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
	goRoutineTest()
}

func goRoutineTest() {
	// messages := make(chan string)
	messages := make(chan string, 5)

	go func() {
		fmt.Println("go routine started!")
		for {
			msg, ok := <-messages
			if !ok {
				return
			}
			fmt.Printf("%v %v\n", time.Now(), msg)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(5000 * time.Millisecond)
	for i := 1; i < 10; i++ {
		fmt.Printf("add ping: %d\n", i)
		messages <- fmt.Sprintf("ping: %d", i)
		// time.Sleep(250 * time.Millisecond)
	}
	close(messages)
	fmt.Println("Generation is done")
	time.Sleep(12 * time.Second)

}
