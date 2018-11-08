package dque

import (
	"log"
	//"github.com/joncrlsn/dque"
)

// Item is the thing we'll be storing in the queue
type Item struct {
	Name string
	Id   int
}

// ItemBuilder creates a new item and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func ItemBuilder() interface{} {
	return &Item{}
}

func ExampleQueue_main() {
	qName := "item-queue"
	qDir := "/tmp"
	segmentSize := 50

	// Create a new queue with segment size of 50
	q, err := NewOrOpen(qName, qDir, segmentSize, ItemBuilder)
	if err != nil {
		log.Fatal("Error creating new dque ", err)
	}

	// Add an item to the queue
	if err := q.Enqueue(&Item{"Joe", 1}); err != nil {
		log.Fatal("Error enqueueing item ", err)
	}

	log.Println("Size should be 1:", q.Size())

	// You can reconsitute the queue from disk at any time
	// as long as you never use the old instance
	q, err = Open(qName, qDir, segmentSize, ItemBuilder)
	if err != nil {
		log.Fatal("Error opening existing dque ", err)
	}

	// Dequeue an item and act on it
	var iface interface{}
	if iface, err = q.Dequeue(); err != nil {
		if err != EMPTY {
			log.Fatal("Error dequeuing item ", err)
		}
	}

	log.Println("Size should be zero:", q.Size())

	// Assert type of the response to an Item pointer so we can work with it
	item, ok := iface.(*Item)
	if !ok {
		log.Fatal("Dequeued object is not an Item pointer")
	}

	doSomething(item)
}

func doSomething(item *Item) {
	log.Println("Dequeued", item)
}
