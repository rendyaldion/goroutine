package goroutine

import "fmt"
import "time"

// Praktek menggunakan goroutine dan channel pada go
// goroutine adalah minithread atau lightweight thread

// cara berkomunikasi antar goroutine atau mengoper data dari suatu goroutine ke goroutine yang lain adalah menggunakan channel 
var itemsChannel = make(chan string)
var cleanedItemsChannel = make(chan string)

func RunChannel()  {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "harta", "kerang"}

	go menyelam(items)
	go membersihkan()
	go menyimpan()

	time.Sleep(500 * time.Millisecond)
}

func menyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil mendapatkan", item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("Berhasil membersihkan", rawItem)
		cleanedItemsChannel <- rawItem
	}
}

func menyimpan() {
	for rawItem := range cleanedItemsChannel {
		fmt.Println("Berhasil menyimpan", rawItem)
	}
}