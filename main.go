package main

import "github.com/rendyaldion/learn-go/goroutine"

// Praktek menggunakan goroutine dan channel pada go
// concurrency adalah dimana kita dapat menjalankan beberapa fungsi secara bersamaan
// concurrency di go menggunakan goroutine
// goroutine seperti minithread atau lightweight thread
func main()  {
	// goroutine practice
	goroutine.RunChannel()
	goroutine.RunSyncWaitGroup()
}