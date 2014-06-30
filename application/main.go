package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func performHeartbeat(now time.Time) {
	key := fmt.Sprintf("hellogo/%s", os.Getenv("HOSTNAME"))
	value := fmt.Sprintf("%s:%s", os.Getenv("HOST_ADDRESS"), os.Getenv("HOST_PORT"))

	log.Printf("Sending heartbeat -- Key: %s, Value: %s", key, value)

	machines := []string{"172.17.0.80:4001"}
	client := etcd.NewClient(machines)
	_, err := client.Set(key, value, uint64(time.Minute))

	if err != nil {
		log.Printf("Hearbeat failed -- Key: %s, Value: %s -- %s", key, value, err)
	}
}

func registerService() {
	performHeartbeat(time.Now())

	tick := time.Tick(45 * time.Second)
	for now := range tick {
		performHeartbeat(now)
	}
}

func startHTTPServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received request for %s", req.URL.Path)
		io.WriteString(w, "Hello Go!\n")
	})

	return http.ListenAndServe(":8000", nil)
}

func main() {
	go registerService()

	err := startHTTPServer()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
