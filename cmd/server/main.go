package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/pborman/getopt"
)

var (
	help  = false
	usage = false

	address = "localhost"
	port    = uint16(80)

	codes = []int{
		500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 520, 521,
		522, 523, 524, 525, 526, 527, 529, 530, 598, 599,
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())

	getopt.BoolVarLong(&help, "help", 'h', "display this help and exit")
	getopt.BoolVarLong(&usage, "usage", 'u', "display this help and exit")

	getopt.StringVarLong(&address, "address", 'a', "bind address", "localhost")
	getopt.Uint16VarLong(&port, "port", 'p', "bind port", "80")

	if err := getopt.Getopt(nil); err != nil {
		log.Panic(err)
	}
}

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := codes[rand.Intn(len(codes))]
	codestr := fmt.Sprintf("%d", code)

	log.Printf("sending code %s to %s\n", codestr, r.RemoteAddr)

	w.WriteHeader(code)

	body := "<html>\n"
	body += "	<title>\n"
	body += "		" + codestr + "\n"
	body += "	</title>\n"
	body += "	<body>\n"
	body += "		" + codestr + "\n"
	body += "	</body>\n"
	body += "</html>\n"

	w.Write([]byte(body))
}

func main() {
	if help || usage {
		getopt.Usage()
		return
	}

	addr := fmt.Sprintf("%s:%d", address, port)
	log.Printf("binding server 5xx to %s\n", addr)
	if err := http.ListenAndServe(addr, server{}); err != nil {
		log.Panic(err)
	}
}
