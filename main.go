// Author: yangzq80@gmail.com
// Date: 2020-09-07
//
package main

import (
	"github.com/yusys-cloud/analog-network/server"
)

func main()  {

	c:=server.ReadConfig()

	server := server.NewServer(c)

	server.Start()

}






