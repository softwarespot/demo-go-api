package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

func main() {
	port := getPort("SERVER_PORT", "10000")

	var requestsCount atomic.Int32
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		currRequestsCount := strconv.Itoa(int(requestsCount.Add(1)))
		json.NewEncoder(w).Encode(map[string]string{
			"msg":      "demo",
			"requests": currRequestsCount,
			"version":  "v3",
		})
		fmt.Println(currRequestsCount, "requests handled")
	})

	fmt.Println("listening on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error: ", err)
	}
}

func getPort(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
