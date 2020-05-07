package main

import (
	"flag"
	"log"
	"net"

	"github.com/zimmerle/dnscom/pkg/dnscom"
)

var ip = flag.String("ip", "127.0.0.1", "IP to listen on.")
var port = flag.Int("port", 53, "Port to listen on.")
var prefixOffset = flag.Int("prefix", 2, "Amount of names to be disposable. \n"+
	"e.g. data.google.com, only data is relevant, \n"+
	"therefore offset will be 2.")
var returnIP = flag.String("return-ip", "8.8.8.8", "ip addr to return in the calls.")

func main() {
	plug := flag.String("plugin", "", "plugin to process the retrieved data.")
	flag.Parse()

	ipNet := net.ParseIP(*ip)
	if ipNet == nil {
		panic("Not a valid ip: " + *ip)
	}

	returnIPNet, _, err := net.ParseCIDR(*returnIP + "/24")
	if err != nil {
		panic(err)
	}

	dnsPlugin, err := dnscom.LoadPlugin(*plug)
	if err != nil {
		panic(err)
	}

	log.Print("Starting dnscom server")
	dnscom.Server(ipNet, returnIPNet, dnsPlugin)

	dnsPlugin.Clean()
}
