package main

import (
    "fmt"
    "net/http"
    "crypto/md5"
    "encoding/base64"
    "github.com/mssola/user_agent"
    "github.com/dchest/uniuri"
)

func randomName(length int) string {
    return uniuri.NewLen(length)
}

func hashName(id string) string {
    h := md5.New()
    h.Write([]byte(id))
    return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    ua := new(user_agent.UserAgent)
    ua.Parse(r.UserAgent())
    id := r.FormValue("id")
    var name string
    if len(id) != 0 {
      name = hashName(id)
    } else {
      name = randomName(10)
    }
    fmt.Fprintf(w, "%s", name)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
