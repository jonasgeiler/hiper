package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"log"
	"os"
)

func main() {
	// Get System Hostname
	hostname, _ := os.Hostname()

	var (
		hcloudApiKey   string
		serverName     string = hostname
		floatingIpName string
	)

	// Define flags (with env fallback)
	flag.StringVar(&hcloudApiKey, "hcloud-api-key", LookupEnvOrString("HCLOUD_API_KEY", hcloudApiKey), "Hetzner Cloud API key (env \"HCLOUD_API_KEY\")")
	flag.StringVar(&serverName, "server-name", LookupEnvOrString("SERVER_NAME", serverName), "Server name (env \"SERVER_NAME\")")
	flag.StringVar(&floatingIpName, "floating-ip-name", LookupEnvOrString("FLOATING_IP_NAME", floatingIpName), "Floating IP name (env \"FLOATING_IP_NAME\")")

	// Parse flags
	flag.Parse()

	// Check for HCloud API key
	if hcloudApiKey == "" {
		log.Fatalf("Hetzner Cloud API Key not specified! Use the \"-hcloud-api-key\" flag or the \"HCLOUD_API_KEY\" environment variable to specify it.")
	}

	// Check for server name
	if serverName == "" {
		log.Fatalf("Server name not specified! Use the \"-server-name\" flag or the \"SERVER_NAME\" environment variable to specify it. Defaults to hostname.")
	}

	// Check for floating IP
	if floatingIpName == "" {
		log.Fatalf("Floating IP name not specified! Use the \"-floating-ip-name\" flag or the \"FLOATING_IP_NAME\" environment variable to specify it.")
	}

	// Initialize HCloud Client
	client := hcloud.NewClient(hcloud.WithToken(hcloudApiKey))

	// Try to get server by name
	server, _, err := client.Server.GetByName(context.Background(), serverName)
	if err != nil {
		// Log request error
		log.Fatalf("Error retrieving server: %s", err)
	}

	if server != nil {
		fmt.Println(fmt.Sprintf("Found server \"%v\".", server.Name))

		// Get floating IP
		ip, _, err := client.FloatingIP.Get(context.Background(), floatingIpName)
		if err != nil {
			// Log request error
			log.Fatalf("Error retrieving floating IP: %s", err)
		}

		if ip != nil {
			fmt.Println(fmt.Sprintf("Found floating IP \"%v\".", ip.Name))

			// Assign floating IP to current server
			_, res, err := client.FloatingIP.Assign(context.Background(), ip, server)
			if err != nil {
				// Log request error
				log.Fatalf("Error assigning floating IP to server: %s", res.Body)
			}

			fmt.Println(fmt.Sprintf("Assigned floating IP \"%v\" to server \"%v\".", ip.Name, server.Name))
		} else {
			// Handle floating IP not found
			fmt.Println(fmt.Sprintf("Unable to find floating IP named \"%v\"!", floatingIpName))
		}

	} else {
		// Handle server not found
		fmt.Println(fmt.Sprintf("Unable to find server named \"%v\"!", serverName))
	}
}

func LookupEnvOrString(env string, defaultVal string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}

	return defaultVal
}
