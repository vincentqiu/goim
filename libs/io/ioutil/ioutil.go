package ioutil

import (
	"bufio"
	log "code.google.com/p/log4go"
)

func ReadAll(rd *bufio.Reader, d []byte) (err error) {
	tl, n, t := len(d), 0, 0
	for {
		if t, err = rd.Read(d[n:]); err != nil {
			return
		}
		if n += t; n == tl {
			break
		} else if n < tl {
			log.Debug("rd.Read() need %d bytes", tl-n)
		} else {
			log.Error("body: readbytes %d > %d", n, tl)
		}
	}
	return
}
