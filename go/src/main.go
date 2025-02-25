package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	// "io"
	"net/http"
	"os"
	"github.com/radjosh/gomm/go/src/dataProcessors"
)

func monsterPtrToJSON(m *dataProcessors.Monster) []byte {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(m)
	// fmt.Println("\nUsing Encoder:\n" + buffer.String())
	return buffer.Bytes()
}

func getMonster(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	monsterName := r.URL.Query().Get("name")
	monsterPtr := new(dataProcessors.Monster)
	*monsterPtr = dataProcessors.Monsters[monsterName]
	payload := monsterPtrToJSON(monsterPtr)
	// io.WriteString(w, payload.String())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func main() {
	http.HandleFunc("/monster", getMonster)
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

/* 
 * run with go run main.go
 * front-end should point to external ip address ie outside of the container
 *
 * frontend and backend containers need to be on the same docker network
 * docker network ls
 * docker network connect <network name> gomm
 * docker network connect <network name> reactmm // or whatever the frontend container is called
 */