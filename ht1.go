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
    Status string
    StatusCode int
    Proto  string
    Header http.Header
    Body   interface{}
}

func ht(v htin, cl *http.Client, enc *json.Encoder, wg *sync.WaitGroup){
    if len(v.Url) > 0 {
        resp, _ := cl.Get(v.Url)
        defer resp.Body.Close()
        v.Body, _ = ioutil.ReadAll(resp.Body)
        v.Header, v.Status, v.StatusCode, v.Proto =
        resp.Header, resp.Status, resp.StatusCode, resp.Proto
        if err := json.Unmarshal(v.Body.([]byte), &(v.Body)); err != nil {
           v.Body = string(v.Body.([]byte))
        }
    }

    enc.Encode(&v)

    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    tr := &http.Transport{}
    cl := &http.Client{Transport: tr}
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    enc.SetEscapeHTML(false)
    for dec.More() {
        var v htin
        dec.Decode(&v)
        wg.Add(1)
        go ht(v, cl, enc, &wg)
    }
    wg.Wait()
}
