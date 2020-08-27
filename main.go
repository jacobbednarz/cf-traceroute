package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/cloudflare/cloudflare-go"
)

func displayMissingParameterNotice(field string) {
	fmt.Printf("missing parameter: %s is required\n\n", field)
	flag.PrintDefaults()
	os.Exit(1)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func supportedTypes() []string {
	return []string{"icmp", "tcp", "udp", "gre", "gre+icmp"}
}

func main() {
	apiKey := flag.String("api-key", "", "Cloudflare API key to use")
	email := flag.String("email", "", "Cloudflare email to use")
	accountID := flag.String("account-id", "", "Cloudflare account ID to use for requests")
	targets := flag.String("targets", "", "Comma delimitered list of targets to run against")
	colos := flag.String("colos", "", "Comma delimitered list of colocations to run the test from")
	debug := flag.Bool("debug", false, "Increase debug verbosity")
	flag.Parse()

	if *apiKey == "" {
		displayMissingParameterNotice("api-key")
	}

	if *email == "" {
		displayMissingParameterNotice("email")
	}

	if *accountID == "" {
		displayMissingParameterNotice("account-id")
	}

	if !contains(supportedTypes(), *packetType) {
		fmt.Printf("invalid packet-type provided: %q. Must be one of %s\n", *packetType, strings.Join(supportedTypes(), ", "))
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *debug {
		fmt.Printf("[DEBUG] api-key:%s\temail:%s\taccount-id:%s\n", *apiKey, *email, *accountID)
	}

	targetsMap := strings.Split(*targets, ",")
	if len(targetsMap) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	coloMap := strings.Split(*colos, ",")

	client, err := cloudflare.New(*apiKey, *email)
	if err != nil {
		log.Fatal(err)
	}

	opts := cloudflare.DiagnosticsTracerouteConfigurationOptions{
		PacketsPerTTL: 10,
		PacketType:    "icmp",
		MaxTTL:        5,
		WaitTime:      1,
	}

	if *debug {
		fmt.Printf("[DEBUG] using traceroute options: %+v\n", opts)
	}

	r, err := client.PerformTraceroute(*accountID, targetsMap, coloMap, opts)
	if err != nil {
		log.Fatalf("failed to perform traceroute: %s", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)

	for _, target := range r {
		fmt.Printf("\n%s\n", target.Target)
		for _, colo := range target.Colos {
			fmt.Printf("  %s\n", strings.ToUpper(colo.Colo.Name))

			if colo.Error != "" {
				fmt.Printf("    %s\n", colo.Error)
				continue
			}

			for _, hop := range colo.Hops {
				for _, node := range hop.Nodes {
					fmt.Fprintf(w, "    %d\t%s (%s - %s)\t%0.2fms\n", hop.PacketsTTL, node.Name, node.IP, node.Asn, node.MeanRttMs)
				}
			}

			w.Flush()
		}
	}
}
