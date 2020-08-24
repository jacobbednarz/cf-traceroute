# cf-traceroute

Perform `traceroute` from the Cloudflare Network to various targets.

NOTE: Pretty rough first cut and depends on
https://github.com/cloudflare/cloudflare-go/pull/520 to work. Once you
have that dependency, this will work.

Tests and functionality will be expanded in the future once that the
above lands.

## Example

```
$ go run main.go \
  -api-key="EXAMPLE" \
  -email="me@example.com" \
  -account-id="f037e56e89293a057740de681ac9abbe" \
  -targets "example.com,google.com,cloudflare.com" 
  -colos "den01,sfo06"

 cloudflare.com
  den01
    1	104.17.176.85 (104.17.176.85 - AS13335)	0.015100ms
  sfo06
    1	104.17.176.85 (104.17.176.85 - AS13335)	0.086600ms

google.com
  sfo06
    1	173.245.61.1 (173.245.61.1 - AS13335)			0.514900ms
    2	NO RESPONSE ( - )					0.000000ms
    3	ae10.cr0-sfo1.ip4.gtt.net (173.205.37.5 - AS3257)	0.431000ms
    4	ae15.cr2-sjc1.ip4.gtt.net (89.149.129.221 - AS3257)	1.646500ms
    5	as15169.sjc10.ip4.gtt.net (199.229.230.134 - *)		2.006300ms
  den01
    1	_gateway (172.68.33.1 - AS13335)			1.074900ms
    2	172.68.32.3 (172.68.32.3 - AS13335)			0.411300ms
    3	108.170.252.209 (108.170.252.209 - AS15169)		1.542100ms
    4	216.239.49.43 (216.239.49.43 - AS15169)			0.431000ms
    5	den02s02-in-f14.1e100.net (172.217.12.14 - AS15169)	0.290800ms
    
example.com
  den01
    1	_gateway (172.68.33.1 - AS13335)				1.049400ms
    2	204.148.254.45 (204.148.254.45 - *)				0.490900ms
    3	204.148.254.6 (204.148.254.6 - *)				1.391000ms
    4	ae-66.core1.dna.edgecastcdn.net (152.199.97.129 - AS15133)	1.051800ms
    5	93.184.216.34 (93.184.216.34 - AS15133)				0.119700ms
  sfo06
    1	173.245.61.1 (173.245.61.1 - AS13335)						0.521600ms
    2	NO RESPONSE ( - )								0.000000ms
    3	ae-21.r05.plalca01.us.bb.gin.ntt.net (131.103.117.73 - AS2914)			1.482100ms
    4	ae-15.r01.snjsca04.us.bb.gin.ntt.net (129.250.5.33 - AS2914)			2.868800ms
    5	ae-0.edgecast-networks.snjsca04.us.bb.gin.ntt.net (129.250.193.134 - AS2914)	2.561900ms
```
