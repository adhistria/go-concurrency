package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = FakeSearch("web")
	Web2   = FakeSearch("web")
	Image1 = FakeSearch("image")
	Image2 = FakeSearch("image")
	Video1 = FakeSearch("video")
	Video2 = FakeSearch("video")
)

type Search func(query string) Result

type Result struct {
	Title string
}

func FakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("result for %s %s\n", kind, query)}
	}
}

// it's almost never timed out because the search is very fast and call to multiple replicas
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// results := First("golang", FakeSearch("replica 1"), FakeSearch("replica 2"))
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}
	return results
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	// return first completed search
	// run in parallel and return the first result
	return <-c
}
