// file create by daihao, time is 2018/7/26 13:03
package main

import (
	"net/url"
	"fmt"
)

func main() {
	rawurl := "http://10.20.1.20:3000/d/-zhYWwOmk/online-txlb?panelId=14&orgId=1&var-UnloadbalancerId=lb-6dh8wa79&var-vip=140.143.252.28&var-lbport=443&var-lanIp=10.131.1.20&var-port=8360&var-InstanceName=apiv1v.main.bjtb.pdtv.it&var-metric=inpkg&from=1532477051465&to=1532650185348"
	u, err := url.Parse(rawurl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.Port())
	u, err = url.ParseRequestURI(rawurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u.User)

	//"0123456789ABCDEF"[c>>4]
}
