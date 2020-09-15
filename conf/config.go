// Author: yangzq80@gmail.com
// Date: 2020-09-07
//
package conf

type Conf struct {
	ApiPort		string		`json:"apiPort,omitempty"`
	Proxies		[]*Proxy	`json:"proxies,omitempty"`
}
// ProxyUnit proxyUnit
type Proxy struct {
	Port            string `json:"port,omitempty"`
	Target         	string `json:"target,omitempty"`
	Desc           	string `json:"desc,omitempty"`
	TimeoutConnect 	int    `json:"timeoutConnect,omitempty"`
	TimeoutWrite   	int    `json:"timeoutWrite,omitempty"`
	Ctl            	*Ctl   `json:"ctl,omitempty"`
}

type Ctl struct {
	In      		*CtlUnit `json:"in"`
	Out     		*CtlUnit `json:"out"`
}

type CtlUnit struct {
	LossRate int `json:"lossRate"`
	DelayMs  int `json:"delayMs"`
}


// CopyFrom copy form
func (c *Ctl) CopyFrom(from *Ctl) {
	c.In.LossRate = from.In.LossRate
	c.In.DelayMs = from.In.DelayMs

	c.Out.LossRate = from.Out.LossRate
	c.Out.DelayMs = from.Out.DelayMs
}