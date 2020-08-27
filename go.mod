module cf-traceroute

go 1.13

require (
	github.com/cloudflare/cloudflare-go v0.13.2-0.20200827221242-38003bd84051
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/text v0.3.3 // indirect
)

replace github.com/cloudflare/cloudflare-go => ../../cloudflare/cloudflare-go
