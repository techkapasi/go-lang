package main

import ("fmt")

func (this IPAddr) String() string {
    var ipaddress string = ""
    for count:=0 ; count < len(this)-1 ; count++ {
        ipaddress += fmt.Sprintf("%v", this[count]) + "."
    }
    ipaddress += fmt.Sprintf("%v", this[len(this)-1])
    return ipaddress
}

type Stringer interface {
    String() string
}

type IPAddr [4]byte

func main () {
    fmt.Println("")
    fmt.Println("Interfaces Mapping:")
    fmt.Println("-------------------")

    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}