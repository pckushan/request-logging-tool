package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	flagName = `parallel`
)

//Worker include a client for the requests
// jobs string channel to input urls
// results string channel to write all responses
type Worker struct {
	client  http.Client
	jobs    <-chan string
	results chan<- string
}

// InitWorker initiate a new worker
// jobs input channel and results output channels are injected to the worker
func InitWorker(jobs, results chan string) *Worker {
	return &Worker{
		client:  http.Client{},
		jobs:    jobs,
		results: results,
	}
}

// start combined worker
// it will make the request for an url and generate a md5 hash for thr response content
// will print the results to the resulting channel
func (w *Worker) start() {
	for j := range w.jobs {
		content, err := w.makeRequest(j)
		if err != nil {
			w.results <- fmt.Sprintf(j + " " + err.Error())
			continue
		}
		outHash := w.getMd5Hash(content)
		w.results <- fmt.Sprintf(j + " " + outHash)
	}
}

func main() {
	// capture flags
	// flag is defined and default value also given
	flagValue := flag.Int(flagName, 10, "number of parallel requests at a time")
	flag.Parse()

	// set argument values captured
	parallelWorkers := *flagValue
	urls := flag.Args()
	// resolve domains
	reAdjustURLs(urls)

	// input and output channels
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	// initiate parallel workers
	for w := 1; w <= parallelWorkers; w++ {
		worker := InitWorker(jobs, results)
		// start initiated worker [non-blocking]
		go worker.start()
	}

	// add urls to the jobs channels [non-blocking]
	go func() {
		for j := 0; j < len(urls); j++ {
			jobs <- urls[j]
		}
		close(jobs)
	}()

	// print from the resulting channel [blocking]
	for a := 1; a <= len(urls); a++ {
		fmt.Println(<-results)
	}
}

// makeRequest takes an URL as an input and send a request
// return an array of bytes and an error if exist
func (w *Worker) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// reAdjustURLs reforms the domain schemas which misses them
func reAdjustURLs(urls []string) []string {
	for i, u := range urls {
		if strings.HasPrefix(u, `http://`) || strings.HasPrefix(u, `https://`) {
			continue
		}
		urls[i] = `http://` + u
	}

	return urls
}

// getMd5Hash generate a md5 hash code for the input array of bytes
func (w *Worker) getMd5Hash(content []byte) string {
	h := md5.New()
	h.Write(content)
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}
