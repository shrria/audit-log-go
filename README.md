# audit-log-go

![example](./staticfiles/example.png)

## Tech Stack

- Elasticsearch
- Kibana

## How to Run

- 1. Start Elasticsearch and Kibana in Docker
     `cd deployment`
     `docker compose up`

- 2. Get Password & Token from Elasticsearch
     `docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic`
     `docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-create-enrollment-token -s kibana`

- 3. Login to Kibana at
     `http://localhost:5601`

- 4. Copy CA Certificate from Elasticsearch to local
     `docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt ../cert`

- 5. Start Go Application (Don't forget to change the password in `elasticsearch.go` file)
     `go build -o audit-log-go`
     `./audit-log-go`

## References

- [Building an Audit Log System for a Go Application](https://medium.com/@alameerashraf/building-an-audit-log-system-for-a-go-application-ce131dc21394)
