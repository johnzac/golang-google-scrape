package googletextresults
import (
"fmt"
"net/http"
"os"
"encoding/json"
"io/ioutil"
"regexp"

)
type Query struct {
	Query string `json:"query"`
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func GetResults(w http.ResponseWriter, request *http.Request) {
    var urlSlice []string
    decoder := json.NewDecoder(request.Body)
    var q Query
    err := decoder.Decode(&q)
    check(err)
    fmt.Println(q.Query)
    googleBaseUrl := "https://www.google.com/search?q="
    googleFinalUrl := googleBaseUrl + q.Query
    client := &http.Client{}
    req,err := http.NewRequest("GET",googleFinalUrl,nil)
    check(err)
    req.Header.Add("user-agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
    resp, err := client.Do(req)
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    googleResponse := string(body)
    r, _ := regexp.Compile("\\<h3 class=\"r\"\\>\\<a href=\"(.*?)\"")
    resultUrls := r.FindAllStringSubmatch(googleResponse,-1)
    for _, resultUrl := range resultUrls {
        fmt.Println(resultUrl[1])
	urlSlice = append(urlSlice, resultUrl[1])
    }
    urlsJson, _ := json.Marshal(urlSlice)
    fmt.Fprintf(w,string(urlsJson))

    }
