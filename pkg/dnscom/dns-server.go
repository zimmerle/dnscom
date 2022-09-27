package dnscom

import (
	"encoding/base32"
	"flag"
	"log"
	"net"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func newRequest(udpConn *net.UDPConn) (*layers.DNS, net.Addr) {
	tmp := make([]byte, 4096)
	_, addr, _ := udpConn.ReadFrom(tmp)

	packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
	dnsPacket := packet.Layer(layers.LayerTypeDNS)
	request, isDNS := dnsPacket.(*layers.DNS)

	if isDNS == false {
		return nil, nil
	}

	return request, addr
}

func anwser(udpConn *net.UDPConn, request *layers.DNS, clientAddr net.Addr, returnIP net.IP) {
	request.QR = true
	request.ANCount = 1
	request.OpCode = layers.DNSOpCodeQuery
	request.AA = true

	for _, question := range request.Questions {
		var dnsAnswer layers.DNSResourceRecord
		if question.Type == layers.DNSTypeAAAA {
			dnsAnswer.IP = net.ParseIP("2600:3c01::f03c:91ff:fe73:5f54")
		} else {
			dnsAnswer.IP = returnIP
		}
		dnsAnswer.Type = question.Type
		dnsAnswer.Name = question.Name
		dnsAnswer.Class = layers.DNSClassIN
		request.Answers = append(request.Answers, dnsAnswer)
	}

	request.ResponseCode = layers.DNSResponseCodeNoErr

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	var err = request.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}
	udpConn.WriteTo(buf.Bytes(), clientAddr)
}

func process(request *layers.DNS, prefixOffset int) (string, string) {
	var requestName []string
	requestName = strings.Split(string(request.Questions[0].Name[:]), ".")
	if len(requestName) <= prefixOffset {
		return "", ""
	}

	return strings.ToUpper(strings.Join(requestName[:len(requestName)-prefixOffset], "")),
		strings.Join(requestName[len(requestName)-prefixOffset:], ".")
}

// Server Main DNS COM Server
func Server(ip net.IP, returnIP net.IP, plug Plugin) {
	port := flag.Lookup("port").Value.(flag.Getter).Get().(int)
	prefixOffset := flag.Lookup("prefix").Value.(flag.Getter).Get().(int)
	address := net.UDPAddr{
		Port: port,
		IP:   ip}

	log.Printf("Listening on: %s at port: %d", ip.String(), port)
	log.Printf("Using consumer plugin: %s", plug.Name())
	log.Printf("Resolving anything to: %s", returnIP.String())
	log.Printf("Dropping %d chunks to find the data.", prefixOffset)

	udpConn, err := net.ListenUDP("udp", &address)
	if err != nil {
		panic(err)
	}

	for {
		request, addr := newRequest(udpConn)
		if request != nil {
			data, res := process(request, prefixOffset)
			padding := (8 - (len(data) % 8)) % 8
			data = data + strings.Repeat("=", padding)
			data2, err := base32.StdEncoding.DecodeString(data)
			if err != nil {
				log.Printf(addr.String()+" Err: %s (Dropped: %s)", data, res)
				plug.Err(data, err)
			} else {
				log.Printf(addr.String()+" %s ", data)
				plug.Ok(strings.Split(addr.String(), ":")[0], string(data2[:]))
			}
			anwser(udpConn, request, addr, returnIP)
		}
	}
}
