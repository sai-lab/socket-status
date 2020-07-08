package functions

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shiro8945/socket-status/lib/status"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	log.SetFlags(0)
	d := status.GetServerStat()

	j, _ := json.Marshal(d)
	buf.Write(j)
	w.Write(buf.Bytes())
}
