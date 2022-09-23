package dnscom

import (
	"errors"
	"plugin"
)

// Plugin methods that needs to be implemented on the plugins
type Plugin interface {
	Init() error
	Clean()
	Err(data string, err error)
	Ok(host string, data string)
	Name() string
}

// LoadPlugin if something went wrong error will be filled.
func LoadPlugin(plug string) (Plugin, error) {
	var dnsPlugin Plugin = nil

	if plug == "" {
		dnsPlugin = &DefaultPlugin{}
		err := dnsPlugin.Init()
		if err != nil {
			return nil, err
		}

		return dnsPlugin, nil
	}

	var err error
	plugSO, err := plugin.Open(plug)
	if err != nil {
		return nil, err
	}

	symPlugin, err := plugSO.Lookup("DNSCOMPlugin")
	if err != nil {
		return nil, err
	}

	var ok bool
	dnsPlugin, ok = symPlugin.(Plugin)
	if !ok {
		return nil, errors.New("failed to load plugin")
	}

	err = dnsPlugin.Init()
	if err != nil {
		return nil, err
	}

	return dnsPlugin, nil
}
