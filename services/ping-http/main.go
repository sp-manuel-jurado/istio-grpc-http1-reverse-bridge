package main

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
)

const (
	addr = ":8888"
)

func main() {
	http.HandleFunc("/sp.rpc.PingService/Ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/sp.rpc.PingService/Ping called!")

		w.Header().Set("Content-Type", "application/grpc+proto")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.WriteHeader(http.StatusOK)

		content, err := proto.Marshal(&empty.Empty{})
		if err != nil {
			panic(err)
		}

		_, err = w.Write(content)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/ping called!")
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
		//_, err := w.Write([]byte("{}"))
		//if err != nil {
		//	panic(err)
		//}
	})

	log.Printf("http server in localhost:" + addr + " started")
	log.Printf("\t- route (http1.1): /ping" + addr)
	log.Printf("\t- route (grpc-http1-reverse-bridge): /sp.rpc.PingService/Ping" + addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
