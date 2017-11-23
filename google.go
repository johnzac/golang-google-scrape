package main
import (
"net/http"
"os"
"github.com/johnzac/googleScrape/googletext"
"fmt"

)

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


func main() {
    http.HandleFunc("/getSearch",googletextresults.GetResults)
    http.ListenAndServe(":8000", nil)
    }

