package main

import (
	"strings"
	xj "github.com/basgys/goxml2json"
	"net/http"
	"flag"
	"log"
	"gopkg.in/tylerb/graceful.v1"
	"time"
	"bytes"
	"io"
	"io/ioutil"
)

var (
	forwarderUrlString = ""
	returnResponseBody = false
)

func main() {

	forwarderUrl := flag.String("forwarder-url", "", "URL to which payload should be forwarded")
	addr := flag.String("addr", "0.0.0.0:8080", "IP/port for the HTTP server")
	returnResponseBodyArg := flag.Bool("return-response-body", false, "Return response body for debug purposes.")
	flag.Parse()

	if *forwarderUrl == "" {
		log.Fatal("The '-forwarder-url' flag is required.")
	} else {
		forwarderUrlString = *forwarderUrl
	}

	returnResponseBody = *returnResponseBodyArg

	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:        *addr,
			ReadTimeout: time.Duration(5) * time.Second,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				json, err := xj.Convert(r.Body)
				if err != nil {
					panic("That's embarrassing...")
				}

				reader := strings.NewReader(json.String())

				request, err := http.NewRequest("POST", forwarderUrlString, reader)
				// TODO: check err

				request.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(request)

				if err != nil {
					log.Fatal(err)
				}

				log.Println(resp.Status)

				if returnResponseBody == true {
					body, err := ioutil.ReadAll(resp.Body)

					if err != nil {
						log.Fatal(err)
					}

					bodyString := string(body)

					log.Println(bodyString)

					buf := &bytes.Buffer{}
					buf.WriteString(bodyString)
					io.Copy(w,buf)
				}
			}),
		},
	}

	log.Println("Server is now listening on ", *addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}