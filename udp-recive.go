package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/ipv4"
	"net"
	"strconv"
)

func main() {
	var nic string

	flag.StringVar(&nic, "i", "", "")
	flag.Parse()

	iface, err := net.InterfaceByName(nic)
	if err != nil {
		panic(err)
	}

	addr, _ := iface.Addrs()
	ip, _, _ := net.ParseCIDR(addr[0].String())
	listenPort := fmt.Sprintf("%s:%d", ip, 10718)
	packet, err := net.ListenPacket("udp4", listenPort)
	if err != nil {
		panic(err)
	}
	defer packet.Close()

	p := ipv4.NewPacketConn(packet)
	if err := p.SetControlMessage(ipv4.FlagInterface, true); err != nil {
		panic(err)
	}

	buffer := make([]byte, 1500)
	for {
		_, cm, addr, err := p.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v:\n", addr)
		fmt.Printf("%s\n", buffer)

		_, portStr, err := net.SplitHostPort(addr.String())
		if err != nil {
			panic(err)
		}

		port, _ := strconv.Atoi(portStr)
		addr = &net.UDPAddr{IP: net.IPv4bcast, Port: port}

		// nilをぶち込まないとエラーになる。。
		// panic: write udp4: invalid argument
		cm.Src = nil
		if _, e := p.WriteTo(buffer, cm, addr); e != nil {
			panic(e)
		}
	}
}
