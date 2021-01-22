# go-dyndns - Yet another Dynamic DNS Tool, this one in GO

Determine remote IPs (IPv4 and IPv6) and adds/updates records to your DNS provider.

## Usage

````shell
go-dyndns --dns provider --host this.is.my.domain.tld
````

## Building

    make build-cli

Or for a compressed binary, using [upx](https://github.com/upx/upx); note this takes longer!

    make dist-cli

## Concept

Use an external service to determine current IP(s)

 * [x] external ipv4 service: http://ip4only.me/api/
 * [x] external ipv6 service: http://ip6only.me/api/

Update the found IP(s) to a DNS provider

 * [x] easy to add new DNS providers
