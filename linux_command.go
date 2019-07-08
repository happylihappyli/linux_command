package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "flag"
    "strings"
)

func httpGet(key string) {
    v := url.Values{}
    v.Add("key",key)
    v.Add("id", "21")
    v.Add("web", "1")
    v.Add("no_topic", "1")
    params := v.Encode()
    
    resp, err := http.Get("https://www.funnyai.com/funnyai/fs_ai_reply.php?"+params)
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    result:=string(body)
    result = strings.Replace(result, "<pre>", "", -1)
    result = strings.Replace(result, "</pre>", "", -1)
  
    fmt.Println(result)
}

func main() {
    key := flag.String("word", "cpu", "a string")
    flag.Parse()
    httpGet(*key)
}
