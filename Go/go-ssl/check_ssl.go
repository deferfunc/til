package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func main() {
	domains := os.Getenv("TARGET_DOMAINS")
	if domains == "" {
		fmt.Println("TARGET_DOMAINS is not set")
		return
	}

	for _, domain := range strings.Split(domains, ",") {
		conn, err := tls.Dial("tcp", net.JoinHostPort(domain, "443"), nil)
		if err != nil {
			fmt.Println("Failed to connect:", err)
			continue
		}
		defer conn.Close()

		fmt.Println("Checking:", domain)

		cert := conn.ConnectionState().PeerCertificates[0]
		expiry := cert.NotAfter
		fmt.Println("SSL certificate expiry:", expiry)

		issuer := cert.Issuer
		fmt.Println("SSL certificate issuer:", issuer)

		daysLeft := expiry.Sub(time.Now()).Hours() / 24
		fmt.Printf("Days left until expiry: %.0f\n", daysLeft)

		if daysLeft < 10 {
			webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
			if webhookURL == "" {
				fmt.Println("SLACK_WEBHOOK_URL is not set")
				continue
			}

			msg := SlackMessage{
				Text: fmt.Sprintf("SSL certificate for %s is about to expire in %.0f days. The issuer is %s.", domain, daysLeft, issuer),
			}
			jsonValue, _ := json.Marshal(msg)

			_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				fmt.Println("Failed to send message to Slack:", err)
			}
		}
	}
}
