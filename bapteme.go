package main

import (
    "fmt"
    "net/http"
    "flag"
    "strings"
    "strconv"
    "encoding/base64"
    "github.com/mssola/user_agent"
    "github.com/dchest/uniuri"
)

var Size int

func prefix(r *http.Request) string {
    ua := new(user_agent.UserAgent)
    ua.Parse(r.UserAgent())

    os := ua.OS()
    if strings.Contains(os, "Linux") {
        return "lin"
    } else if strings.Contains(os, "Windows") {
        return "win"
    } else {
        return "srv"
    }
}

func randomName(length int) string {
    return uniuri.NewLen(length)
}

func hashName(id string) string {
//    h := md5.New()
//    h.Write([]byte(id))
//    return base64.URLEncoding.EncodeToString(h.Sum(nil))
    return base64.URLEncoding.EncodeToString([]byte(id))
}

func handler(w http.ResponseWriter, r *http.Request, size int) {
    id := r.FormValue("id")

    tmpsize := r.FormValue("size")

    if len(tmpsize) != 0 {
      s, err := strconv.Atoi(tmpsize)
      size = s
      if err != nil {
        fmt.Println(err)
        return
      }
    }

    prefix := prefix(r)
    var suffix string
    if len(id) != 0 {
      suffix = hashName(id)[0:size-3]
    } else {
      suffix = randomName(size-3)
    }
    name := strings.Join([]string{ prefix, suffix},"")
    fmt.Fprintf(w, "%s", name)
}

func main() {
    port := flag.Int("port", 8080, "Port to use")
    address := flag.String("address", "", "Address to bind")
    size := flag.Int("size", 10, "Default final hostname size")

    flag.Parse()

    socket := fmt.Sprint(*address, ":", *port)
    fmt.Printf("Bind to %s", socket)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
              handler(w, r, *size)
       })
    http.ListenAndServe(socket , nil)
}
