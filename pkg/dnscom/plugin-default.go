package dnscom

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// DefaultPlugin holds the Default plugin data.
type DefaultPlugin struct {
	rdb *redis.Client
}

type DNSQueryData struct {
	Host string
	Data string
	Ok   bool
}

func (j DNSQueryData) MarshalBinary() ([]byte, error) {
	return json.Marshal(j)
}

// Err call in case of an error.
func (p *DefaultPlugin) Err(a string, b error) {
	fmt.Println("E> " + a + ": " + b.Error())
	dnsqd := DNSQueryData{
		Host: "-",
		Data: a,
		Ok:   false,
	}

	payload, err := json.Marshal(dnsqd)
	if err != nil {
		fmt.Println("E2> " + a + ": " + err.Error())
		return
	}

	if err := p.rdb.Publish(context.Background(), "testing", payload).Err(); err != nil {
		panic(err)
	}
}

// Ok call in case of a function
func (p *DefaultPlugin) Ok(host string, a string) {
	fmt.Println(host + "!> " + a)

	dnsqd := DNSQueryData{
		Host: host,
		Data: a,
		Ok:   true,
	}

	payload, err := json.Marshal(dnsqd)
	if err != nil {
		fmt.Println("E2> " + a + ": " + err.Error())
		return
	}

	if err := p.rdb.Set(context.Background(), a+"+dns", dnsqd, 0).Err(); err != nil {
		panic(err)
	}

	if err := p.rdb.Publish(context.Background(), "testing", payload).Err(); err != nil {
		panic(err)
	}

}

// Init init the plugin.
func (p *DefaultPlugin) Init() error {
	p.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return nil
}

// Clean cleanup the plugin
func (p *DefaultPlugin) Clean() {
}

func (p *DefaultPlugin) Name() string {
	return "Default"
}

var DNSCOMPlugin DefaultPlugin
