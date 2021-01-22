package providers_ip_default

import (
	"bytes"
	"fmt"
	"github.com/AubreyHewes/go-dyndns/v1/log"
	"github.com/AubreyHewes/go-dyndns/v1/providers/ip/types"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
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

func (d *IPProvider) GetIPs() ([]types.IPResult, error) {
	ipv4, err := d.GetIPv4()
	if err != nil {
		log.Verbosef("IP Provider [default] GetIPv4 error: %s", err)
	}
	ipv6, err := d.GetIPv6()
	if err != nil {
		log.Verbosef("IP Provider [default] GetIPv6 error: %s", err)
	}

	var ips = []types.IPResult{}

	if ipv4 != nil && ipv6 != nil && ipv4.String() == ipv6.String() {
		ips = append(ips, types.IPResult{Type: "A", IP: ipv4})
		return ips, nil
	}

	if ipv4 != nil {
		ips = append(ips, types.IPResult{Type: "A", IP: ipv4})
	}
	if ipv6 != nil {
		ips = append(ips, types.IPResult{Type: "AAAA", IP: ipv6})
	}
	return ips, nil
}

func (d *IPProvider) GetIPv4() (net.IP, error) {
	var url = "http://ip4only.me/api/"
	log.Verbosef("IP Provider [default] GetIPv4 url: %s", url)
	res, err := d.client.Get(url)
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
	log.Verbosef("IP Provider [default] GetIPv4 output: %s", body)

	var parts = strings.Split(string(body), ",")
	if parts[0] != "IPv4" {
		return nil, fmt.Errorf("%s", "output is not a IPv4 result")
	}

	err = ipip.UnmarshalText(bytes.TrimSpace([]byte(parts[1])))
	if err != nil {
		return nil, err
	}
	return ipip, nil
}

func (d *IPProvider) GetIPv6() (net.IP, error) {
	var url = "http://ip6only.me/api/"
	log.Verbosef("IP Provider [default] GetIPv6 url: %s", url)
	res, err := d.client.Get(url)
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
	log.Verbosef("IP Provider [default] GetIPv6 output: %s", body)

	var parts = strings.Split(string(body), ",")
	if parts[0] != "IPv6" {
		return nil, fmt.Errorf("%s", "output is not a IPv6 result")
	}

	err = ipip.UnmarshalText(bytes.TrimSpace([]byte(parts[1])))
	if err != nil {
		return nil, err
	}
	return ipip, nil
}
