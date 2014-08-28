package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/mssola/user_agent"
	"github.com/op/go-logging"
	"net/http"
	"strconv"
	"strings"
)

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

  var err error
  var size int
	formSize := r.FormValue("size")

	if formSize != "" {
		size , err = strconv.Atoi(formSize)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
    size = *flagSize
  }

	var name string

	pre := r.FormValue("prefix")
	if len(pre) != 0 {
		name = pre
	} else {
		name = prefix(r)
	}
	log.Debug("Prefix = %s", name)

	instance := r.FormValue("instance")

	name = strings.Join([]string{name, instance}, "")
	if len(name) >= size {
		http.Error(w, "name too long", 500)
		return
	}

	var suffix string
	if len(id) != 0 {
		ts := hashName(id)
		if len(ts) >= size-len(name) {
			suffix = ts[0 : size-len(name)]
		} else {
			suffix = ts
		}
	} else {
		suffix = randomName(size - len(name))
	}
	log.Debug("Suffix = %s", suffix)
	name = strings.Join([]string{name, suffix}, "")

	fmt.Fprintf(w, "%s", name)
}

var (
  flagBind = flag.String("bind", "", "Address to bind. Format IP:PORT")
  flagSize = flag.Int("size", 10, "Default final hostname size")
  flagDebug = flag.Bool("d", false, "turn on debug info")
)

var log = logging.MustGetLogger("bapteme")

func main() {
	flag.Parse()
	var format = logging.MustStringFormatter("%{level} %{message}")
	logging.SetFormatter(format)
	if *flagDebug {
		logging.SetLevel(logging.DEBUG, "bapteme")
	} else {
		logging.SetLevel(logging.INFO, "bapteme")
	}

	log.Info("Bind to %s", *flagBind)

	http.HandleFunc("/", handler)
	//    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//              handler(w, r, *size)
	//       })
	http.ListenAndServe(*flagBind, nil)
}
