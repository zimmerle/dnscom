![DNSCON](dnscom.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/zimmerle/dnscom)](https://goreportcard.com/report/github.com/zimmerle/dnscom)

DNSCON is an uncomplicated utility to support communication over DNS Queries.


## Background

It is possible to exchange a small portion of data through DNS queries. This method is mainly used in a scenario where there is no other alternative to reach the internet or outside world. Usually, this limitation is encountered due to firewall restrictions, which tend to be eased for DNS queries.

For the communication to flow, there is a need for two figures: (a) the **server**, (b) the **client**. The **client** is pretty much any application that is capable of resolving names (e.i. perform DNS queries). The **server**, in the order hand, needs to be crafted for that purpose. The **server** is including in this package.

#### Making the client reaching the server

There are multiples ways to make the **client** performing requests on your **server**; the most straightforward one is to use a DNS tool such us: [dnsq](https://cr.yp.to/djbdns/debugging.html), [drill](https://www.nlnetlabs.nl/projects/ldns/about/), or [nslookup](https://dougbarton.us/DNS/bind-users-FAQ.html) to query your server.

Preferable and bolder option to drive the **clients** towards your **server** is to delegate a subdomain to be under your server's responsibility. That is preferable as it goes evenly on the producer machine (**client**).

## Server

This package provides an elementary but expandable server that is capable of loading plugins to process the queries payload. Look at a plugin example here: (plugins example)[plugins/example]. If no plugin is loaded, the server will dump the query payload in the stdout. As illustrated in the example below.

![DNSCON](asciicast.gif)

#### Installation

```
go install github.com/zimmerle/dnscom/cmd/dnscom-server
```

#### Command line Help
```
  -ip string
        IP to listen on. (default "127.0.0.1")
  -plugin string
        plugin to process the retrieved data.
  -port int
        Port to listen on. (default 53)
  -prefix int
        Amount of names to be disposable. 
        e.g. data.google.com, only data is relevant, 
        therefore offset will be 2. (default 2)
  -return-ip string
        ip addr to return in the calls. (default "8.8.8.8")
```

## Encoder

This package also includes an Encoder, that is pretty much a utility to transform human-readable data to a format that fits a domain/subdomain, hence, queriable via DNS.

#### Installation

```
go install github.com/zimmerle/dnscom/cmd/dnscom-encode
```
