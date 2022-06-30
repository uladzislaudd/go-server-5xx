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
	address = "localhost"
	port    = uint16(80)

	codes = []int{500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 520, 521, 522, 523, 524, 525, 526, 527, 529, 530, 598, 599}
)

func init() {
	rand.Seed(time.Now().UnixNano())

	getopt.StringVarLong(&address, "address", 'a', "address", "localhost")
	getopt.Uint16VarLong(&port, "port", 'p', "port", "80")
	if err := getopt.Getopt(nil); err != nil {
		panic(err)
	}
}

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := codes[rand.Intn(len(codes))]
	codestr := fmt.Sprintf("%d", code)
	log.Printf("sending code %s to %s\n", codestr, r.RemoteAddr)
	w.WriteHeader(code)
	body := "<html>\n" + "	<title>\n" + codestr + "	</title>\n" + "	<body>\n" + codestr + "	</body>\n" + "</html>\n"
	w.Write([]byte(body))
}

func main() {
	addr := fmt.Sprintf("%s:%d", address, port)
	log.Printf("staring 5** server on %s\n", addr)
	if err := http.ListenAndServe(addr, server{}); err != nil {
		log.Panic(err)
	}
}
