package log

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lbb4511/wechat/constants"
)

func writeLog(r *http.Request, t time.Time, match string, pattern string) {

	if constants.LogLevel != "prod" {

		d := time.Now().Sub(t)

		l := fmt.Sprintf("[ACCESS] | % -10s | % -40s | % -16s | % -10s | % -40s |", r.Method, r.URL.Path, d.String(), match, pattern)

		log.Println(l)
	}
}
