package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func ping(host string) {
	cmd := exec.Command("ping", host, "-c 5")
	fmt.Println(cmd.Args)
	out, _ := cmd.CombinedOutput()
	if strings.Contains(string(out), "5 packets received") {
		arr := strings.Split(string(out), "\n")
		fmt.Println(arr[len(arr)-3] + "\n" + arr[len(arr)-2])
	} else {
		fmt.Println(string(out))
	}

	fmt.Println("")
}

func curl(host string) {
	cmd := exec.Command("curl", "https://"+host, "-I")
	fmt.Println(cmd.Args)
	out, _ := cmd.CombinedOutput()
	if strings.Contains(string(out), "200 OK") {
		arr := strings.Split(string(out), "\n")
		fmt.Println(arr[0])
	} else {
		fmt.Println(string(out))
	}
	fmt.Println("")
}

func nc(host string, port int) {
	cmd := exec.Command("nc", host, strconv.Itoa(port), "-vz")
	fmt.Println(cmd.Args)
	out, _ := cmd.CombinedOutput()
	if strings.Contains(string(out), "succeeded!") {
		arr := strings.Split(string(out), "\n")
		fmt.Println(arr[len(arr)-2])
	} else {
		fmt.Println(string(out))
	}
	fmt.Println("")
}

func traceroute(host string) {
	cmd := exec.Command("traceroute", host)
	fmt.Println(cmd.Args)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))
}

func openssl(host string) {
	cmd := exec.Command("openssl", "s_client", "-connect", host+":443")
	fmt.Println(cmd.Args)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))
}

func main() {
	flag.Parse()
	host := flag.Arg(0)

	ping(host)
	//traceroute(host)
	nc(host, 80)
	nc(host, 443)
	curl(host)
	openssl(host)
}
