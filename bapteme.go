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
    "github.com/op/go-logging"
)

var Size int

func prefix(r *http.Request) string {
    ua := new(user_agent.UserAgent)
    ua.Parse(r.UserAgent())

    os := ua.OS()
    var ret string
    if strings.Contains(os, "Linux") {
        ret = "lin"
    } else if strings.Contains(os, "Windows") {
        ret = "win"
    } else {
        ret = "srv"
    }
    return ret
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

func handler(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")

    tmpsize := r.FormValue("size")

    if len(tmpsize) != 0 {
      s, err := strconv.Atoi(tmpsize)
      *size = s
      if err != nil {
        fmt.Println(err)
        return
      }
    }

    prefix := prefix(r)
    log.Debug("Prefix = %s", prefix)
    var suffix string
    if len(id) != 0 {
      ts := hashName(id)
      if len(ts) >= *size-3 {
        suffix = ts[0:*size-3]
      } else {
        suffix = ts
      }
    } else {
      suffix = randomName(*size-3)
    }
    log.Debug("Suffix = %s", suffix)

    name := strings.Join([]string{ prefix, suffix},"")
    fmt.Fprintf(w, "%s", name)
}

var port    = flag.Int("port", 8080, "Port to use")
var address = flag.String("address", "", "Address to bind")
var size    = flag.Int("size", 10, "Default final hostname size")
var debug   = flag.Bool("d", false,"turn on debug info")

var log = logging.MustGetLogger("bapteme")

func main() {
    flag.Parse()
    var format = logging.MustStringFormatter("%{level} %{message}")
    logging.SetFormatter(format)
    if *debug {
      logging.SetLevel(logging.DEBUG, "bapteme")
    } else {
      logging.SetLevel(logging.INFO, "bapteme")
    }

    socket := fmt.Sprint(*address, ":", *port)
    log.Info("Bind to %s", socket)

    http.HandleFunc("/", handler)
//    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//              handler(w, r, *size)
//       })
    http.ListenAndServe(socket , nil)
}
