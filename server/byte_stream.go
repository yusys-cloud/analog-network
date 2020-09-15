// Author: yangzq80@gmail.com
// Date: 2020-09-09
//
package server

import (
	"github.com/yusys-cloud/analog-network/conf"
	"io"
	"log"
	"time"
)

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func (proxyServer *ProxyServer) CtrlCopyBuffer(dst io.Writer, src io.Reader,cu *conf.CtlUnit, buf []byte) (written int64, err error) {
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}

	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			//Loss rate
			if proxyServer.rnd.Intn(100) < cu.LossRate {
				log.Println("Loss byte ---> ",string(buf))
				continue
			}
			//Delay milliseconds
			if cu.DelayMs > 0 {
				time.Sleep(time.Millisecond * time.Duration(cu.DelayMs))
			}
			//write
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}