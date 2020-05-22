package ifconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// Config is used to configure the creation of the IPProvider
type Config struct {
}

func NewDefaultConfig() *Config {
	return &Config{}
}

// IPProvider is an implementation of the providers.ip.Provider interface
type IPProvider struct {
	config *Config
	client http.Client
}

func NewIPProvider() (*IPProvider, error) {
	config := NewDefaultConfig()
	return NewIPProviderConfig(config)
}

func NewIPProviderConfig(config *Config) (*IPProvider, error) {
	client := http.Client{}
	return &IPProvider{config: config, client: client}, nil
}

func (d *IPProvider) GetIP() (net.IP, error) {
	res, err := d.client.Get("https://ifconfig.co/ip")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var ipip net.IP
	err = ipip.UnmarshalText(bytes.TrimSpace(body))
	if err != nil {
		return nil, err
	}
	return ipip, nil
}
