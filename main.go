package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"sync"
)

var (
	// a lock that protects the queue
	mu sync.Mutex

	// a queue is represented by its start and end position
	start, end int
	key, iv    []byte
	block      cipher.Block
)

func main() {
	key = make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	var err error
	block, err = aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
}

func getQueueTotal() int {
	mu.Lock()
	defer mu.Unlock()
	return end - start
}

func getQueueStart() int {
	mu.Lock()
	defer mu.Unlock()
	return start
}

func getQueueEnd() int {
	mu.Lock()
	defer mu.Unlock()
	return end
}

func getPosition(hash string) int {
	return 0
}

// signTicket issues a ticket for a client to hold, the ticket is generated
// by encrypting position and deviceId, which will be decrypted later to aquire
// the ticket's position and its deviceId for verification use.
func signTicket(position, deviceId int) int {
	return 0
}

// check ticket's position by decrypting its position number and its deviceId
// once the check is passed, the waitingroom generates a token and pass it to
// the client for its later use, it returns a null string when verification is
// not passed or something went wrong.
func tryGetToken() (token string) {
	return ""
}
