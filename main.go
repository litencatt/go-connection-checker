package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func ping(host string) {
	out, _ := exec.Command("ping", host, "-c 5").Output()
	fmt.Println("ping -c 5 " + host)
	fmt.Println(string(out))
}

func curl(host string) {
	out, _ := exec.Command("curl", "https://"+host, "-I").Output()
	fmt.Println("curl -I https://" + host)
	if strings.Contains(string(out), "200 OK") {
		arr := strings.Split(string(out), "\n")
		fmt.Println(arr[0])
	} else {
		fmt.Println(string(out))
	}
	fmt.Println("")
}

func nc(host string, port int) {
	out, _ := exec.Command("nc", host, strconv.Itoa(port), "-vz").CombinedOutput()
	fmt.Println("nc -vz " + host + strconv.Itoa(port))
	if strings.Contains(string(out), "succeeded!") {
		arr := strings.Split(string(out), "\n")
		fmt.Println(arr[len(arr)-2])
	} else {
		fmt.Println(string(out))
	}
	fmt.Println("")
}

func traceroute(host string) {
	out, _ := exec.Command("traceroute", host).Output()
	fmt.Println("traceroute " + host)
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
}
