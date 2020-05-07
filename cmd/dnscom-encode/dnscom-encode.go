package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	dots := flag.Int("subdomain", 0, "subdomain size having N characters.")
	timestamp := flag.Bool("time", true, "Add time stamp.")
	domain := flag.String("domain", "google.com", "Domain.")
	flag.Parse()

	now := time.Now()

	aa := strings.Join(flag.CommandLine.Args(), " ")
	encoded := base32.StdEncoding.EncodeToString([]byte(aa))

	if *dots > 0 {
		for idx := len(encoded) - *dots; idx > 0; idx = idx - *dots {
			encoded = encoded[:idx] + "." + encoded[idx:]
		}
	}

	if *timestamp {
		encoded = encoded + "." + strconv.FormatInt(now.Unix(), 10)
	}

	encoded = strings.ToLower(encoded) + "." + *domain

	fmt.Println(encoded)
}
