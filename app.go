package main

import (
    "fmt"
    "net/http"
    "regexp"
    "io/ioutil"
    "strings"
)


func handler(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path;

	resp, err := http.Get("http://www.thenewslens.com/" + path)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body);

	// http://image.thenewslens.com/wp-content/uploads/2015/09/徵人-banner_徵人.png
	// replace to
	// http://i.imgur.com/lPpksGP.jpg
	patten := "http://image.thenewslens.com/wp-content/uploads/2015/09/徵人-banner_徵人.png"
	new_string := "http://i.imgur.com/lPpksGP.jpg"
	content := string(body)
	content = strings.Replace(content, patten, new_string, -1)

	fmt.Fprintf(w, content )
}

func main() {

	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))


    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}