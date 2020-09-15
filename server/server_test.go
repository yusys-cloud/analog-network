// Author: yangzq80@gmail.com
// Date: 2020-09-07
//
package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStr(t *testing.T)  {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

func test1(p *int )  {
	for {
		time.Sleep(time.Second)
		fmt.Println(*p)
	}
}
