package main

func main() {
	// Log an audit event
	logAuditEvent("John Doe", "password_change", "Password Settings", "Changed password for security reasons.")

	// Send the audit log to Elasticsearch
	sendToElasticsearch()
}
