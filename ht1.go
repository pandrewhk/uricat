package main

//import "fmt"
//import "bufio"
import "net/http"
import "encoding/json"
import "os"
import "io/ioutil"
//import "log"
//import "time"
import "sync"

type htin struct {
    Url    string
    Header http.Header
    Body   interface{}
}

func ht(v htin, enc *json.Encoder, wg *sync.WaitGroup){
    if len(v.Url) > 0 {
        resp, _ := http.Get(v.Url)
        defer resp.Body.Close()
        v.Body, _ = ioutil.ReadAll(resp.Body)
        v.Header = resp.Header
//        v.Body = string(v.Body.([]byte))
        if err := json.Unmarshal(v.Body.([]byte), &(v.Body)); err != nil {
           v.Body = string(v.Body.([]byte))
        }
    }

//    log.Println(v.Body)
    enc.Encode(&v)

    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    enc.SetEscapeHTML(false)
    for dec.More() {
        var v htin
        dec.Decode(&v)
        wg.Add(1)
        go ht(v, enc, &wg)
    }
    wg.Wait()
}
