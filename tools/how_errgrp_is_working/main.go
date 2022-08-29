package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "golang.org/x/sync/errgroup"
)


func main() {
    start := time.Now()

    zenChan := make(chan *http.Response, 1)
    pingChan := make(chan *http.Response, 1)
    helloChan := make(chan *http.Response, 1)

    errGrp, _ := errgroup.WithContext(context.Background())
    errGrp.Go(func() error { return GetAsync("http://localhost:8080/zen.json", zenChan) })
    errGrp.Go(func() error { return GetAsync("http://localhost:8080/ping.json", pingChan) })
    errGrp.Go(func() error { return GetAsync("http://localhost:8080/hello-world.json", helloChan) })

    err := errGrp.Wait()
    if err != nil {
        fmt.Println(err)
        fmt.Println("Error in group async requests")
        os.Exit(1)
    }

    zenResp := <-zenChan
    defer zenResp.Body.Close()
    bytes, _ := ioutil.ReadAll(zenResp.Body)
    fmt.Println(string(bytes))

    pingResp := <-pingChan
    defer pingResp.Body.Close()
    bytes, _ = ioutil.ReadAll(pingResp.Body)
    fmt.Println(string(bytes))

    helloResp := <-helloChan
    defer pingResp.Body.Close()
    bytes, _ = ioutil.ReadAll(helloResp.Body)
    fmt.Println(string(bytes))

    end := time.Now()

    fmt.Printf("Requests processed after %v seconds\n", end.Sub(start).Seconds())
}

func GetAsync(url string, rc chan *http.Response) error {
    resp, err := http.Get(url)
    if err == nil {
        rc <- resp
    }

    return err
}