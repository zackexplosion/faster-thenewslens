package main

import (
    "fmt"
    "net/http"
    "regexp"
    "io/ioutil"
    "strings"
    "os"
)


var re = regexp.MustCompile(`(?m)https?:\/\/[^\s)"]+`)

func handler(w http.ResponseWriter, r *http.Request) {
    var path = r.URL.Path;

    resp, err := http.Get("http://www.thenewslens.com/" + path)
    if err != nil {
    	// handle error
        return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body);


    content := string(body)



    // *****************************************************************************
    // http://image.thenewslens.com/wp-content/uploads/2015/09/徵人-banner_徵人.png
    //
    // replace to
    //
    // http://i.imgur.com/lPpksGP.jpg
    // *****************************************************************************

    patten := "http://image.thenewslens.com/wp-content/uploads/2015/09/徵人-banner_徵人.png"
    new_string := "http://i.imgur.com/lPpksGP.jpg"
    content = strings.Replace(content, patten, new_string, -1)


    // *****************************************************************************
    // search all url
    // re := regexp.MustCompile(`(https?:\/\/[^\s)"]+)`)
    result := re.FindAllString(content, -1)
    for _, v := range result {
        matched := string(v)
        i2 := strings.Index(matched, "wp.com")
        if i2 != -1 {
            new_string := strings.Split(v, "?")[0]
            new_string = new_string + "?w=640"
            content = strings.Replace(content, v, new_string, -1)
        }
    }

    fmt.Fprintf(w, content )
}

func main() {

    http.HandleFunc("/", handler)

    // http.ListenAndServe(":8080", nil)
    http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}