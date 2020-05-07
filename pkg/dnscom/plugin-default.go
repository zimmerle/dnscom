package dnscom

import (
	"fmt"
)

// DefaultPlugin holds the Default plugin data.
type DefaultPlugin struct {
}

// Err call in case of an error.
func (p *DefaultPlugin) Err(a string, b error) {
	fmt.Println("E> " + a + ": " + b.Error())
}

// Ok call in case of a function
func (p *DefaultPlugin) Ok(host string, a string) {
	fmt.Println(host + "> " + a)
}

// Init init the plugin.
func (p *DefaultPlugin) Init() error {
	return nil
}

// Clean cleanup the plugin
func (p *DefaultPlugin) Clean() {
}

func (p *DefaultPlugin) Name() string {
	return "Default"
}

var DNSCOMPlugin DefaultPlugin
