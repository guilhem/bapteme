package main

import (
    "fmt"
    "net/http"
    "flag"
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
    port := flag.Int("port", 8080, "Port to use")
    address := flag.String("address", "", "Address to bind")

    flag.Parse()

    socket := fmt.Sprint(*address, ":", *port)
    fmt.Printf("Bind to %s", socket)
    http.HandleFunc("/", handler)
    http.ListenAndServe(socket , nil)
}
