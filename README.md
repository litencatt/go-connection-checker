## Usage

```
$ go run main.go www.google.com
ping -c 5 www.google.com
PING www.google.com (216.58.220.228): 56 data bytes
64 bytes from 216.58.220.228: icmp_seq=0 ttl=53 time=25.156 ms
64 bytes from 216.58.220.228: icmp_seq=1 ttl=53 time=24.689 ms
64 bytes from 216.58.220.228: icmp_seq=2 ttl=53 time=27.072 ms
64 bytes from 216.58.220.228: icmp_seq=3 ttl=53 time=24.113 ms
64 bytes from 216.58.220.228: icmp_seq=4 ttl=53 time=31.942 ms

--- www.google.com ping statistics ---
5 packets transmitted, 5 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 24.113/26.594/31.942/2.852 ms

nc -vz www.google.com80
Connection to www.google.com port 80 [tcp/http] succeeded!

nc -vz www.google.com443
Connection to www.google.com port 443 [tcp/https] succeeded!

curl -I https://www.google.com
HTTP/1.1 200 OK
```
