package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mattn/go-pipeline"
)

func ping(host string) {
	c1 := []string{"ping", host, "-c 5"}
	c2 := []string{"grep", "received"}
	execCommand(c1, c2)
}

func curl(host string) {
	c1 := []string{"curl", "https://" + host, "-s", "-I"}
	c2 := []string{"grep", "200 OK"}
	execCommand(c1, c2)
}

func nc(host string, port int) {
	c1 := []string{"nc", host, strconv.Itoa(port), "-vz"}
	c2 := []string{"grep", "succeeded!"}
	execCommand(c1, c2)
}

func traceroute(host string) {
	c1 := []string{"traceroute", host}
	execCommand(c1)
}

func openssl(host string) {
	c1 := []string{"openssl", "s_client", "-connect", host + ":443"}
	execCommand(c1)
}

func showCommand(c ...[]string) {
	showCommands := ""
	for _, cmd := range c {
		if showCommands == "" {
			showCommands += strings.Join(cmd, " ")
		} else {
			showCommands += " | " + strings.Join(cmd, " ")
		}
	}
	fmt.Println(showCommands)
}

func execCommand(c ...[]string) {
	showCommand(c...)

	out, err := pipeline.Output(c...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func main() {
	flag.Parse()
	host := flag.Arg(0)

	fmt.Println("Check L3 layer connection.")
	fmt.Println("")
	ping(host)
	//traceroute(host)

	fmt.Println("Check L4-7 layer connection.")
	nc(host, 80)
	curl(host)

	nc(host, 443)
	openssl(host)
}
