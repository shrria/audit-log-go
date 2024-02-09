package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

func sendToElasticsearch() {
	elasticEndpoint := "https://localhost:9200"
	elasticUsername := "elastic"
	elasticPassword := "d8mcclbOQji2VplwaTDU"

	if elasticEndpoint == "" || elasticUsername == "" || elasticPassword == "" {
		log.Fatalf("Error reading the environment variables")
	}

	cert, err := os.ReadFile("../cert/http_ca.crt")
	if err != nil {
		log.Fatalf("Error reading the certificate: %s", err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticEndpoint,
		},
		Username: elasticUsername,
		Password: elasticPassword,
		CACert:   cert,
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	file, err := os.Open("audit.log")
	if err != nil {
		log.Fatalf("Error opening audit.log: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry AuditLogEntry
		if err := json.Unmarshal([]byte(scanner.Text()), &entry); err != nil {
			log.Fatalf("Error parsing JSON: %s", err)
		}

		body, err := json.Marshal(entry)
		if err != nil {
			log.Fatalf("Error marshaling entry to JSON: %s", err)
		}

		req := esapi.IndexRequest{
			Index:        "audit_logs",
			DocumentType: "_doc",
			DocumentID:   "",
			Body:         strings.NewReader(string(body)),
			Refresh:      "true",
		}

		resp, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error sending entry to Elasticsearch: %s", err)
		}
		defer resp.Body.Close()

		log.Printf("Response: %s", resp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading audit.log: %s", err)
	}
}
