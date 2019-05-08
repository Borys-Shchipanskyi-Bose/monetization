package main

import (
	"fmt"
	"time"
)

func main() {
	// http.HandleFunc("/", handlers.Hello)

	// log.Fatal(http.ListenAndServe(":8080", nil))
	goRoutineTest()
}

func goRoutineTest() {
	messages := make(chan string, 10)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			msg, ok := <-messages
			if !ok {
				return
			}
			fmt.Printf("%v %v\n", time.Now(), msg)
		}
	}()

	for i := 0; i < 10; i++ {
		messages <- fmt.Sprintf("ping: %d", i)
	}
	close(messages)
	fmt.Println("Generation is done")
	time.Sleep(12 * time.Second)

}
