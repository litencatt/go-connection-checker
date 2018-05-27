package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func ping(host string) {
	execCommand(
		[]string{"ping", host, "-c 5"},
		"5 packets received",
	)

}

func curl(host string) {
	execCommand(
		[]string{"curl", "https://" + host, "-s", "-I"},
		"",
	)
}

func nc(host string, port int) {
	execCommand(
		[]string{"nc", host, strconv.Itoa(port), "-vz"},
		"succeeded!",
	)
}

func traceroute(host string) {
	execCommand([]string{"traceroute", host}, "")
}

func openssl(host string) {
	execCommand([]string{"openssl", "s_client", "-connect", host + ":443"}, "")
}

func execCommand(command []string, check string) {
	cmd := exec.Command(command[0], command[1:]...)
	fmt.Println(cmd.Args)

	out, _ := cmd.CombinedOutput()
	if check != "" && strings.Contains(string(out), check) {
		arr := strings.Split(string(out), "\n")
		fmt.Println(arr[len(arr)-2])
	} else {
		fmt.Println(string(out))
	}
	fmt.Println("")
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
