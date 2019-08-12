
package main



import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/codegangsta/negroni"
)

type Fed_json struct{
  Message string `json:"message"`
}

func main() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

        message := return_message()

        json_message, err := json.Marshal(message)
        if err != nil {
            log.Println(err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }


        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, string(json_message))
    })
    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":"+os.Args[1])
}

func return_message() *Fed_json{
  return &Fed_json{
    Message : "Hello World!!",
  }
}
