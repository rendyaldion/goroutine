package goroutine

import "fmt"
import "sync"

func RunSyncWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printText("Salam", &wg)

	wg.Add(1)
	go printText("Halo", &wg)

	wg.Wait()
}

func printText(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(text); i++ {
		fmt.Println(text)
	}
}