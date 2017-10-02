package globals

const (
	NAMESPACE        string = "com.go-micro-tech-talk"
	FLIGHT_SRV       string = "flight"
	INDEXER_SRV      string = "indexer"
	ORCHESTRATOR_SRV string = "orchestrator"

	INDEX_FLIGT string = "index-flight-es"
	TYPE_FLIGHT string = "flight"

	INDEX_FLIGHT_TOPIC string = NAMESPACE + ".topic.index.flights"

	INDEX_FLIGHT_QUEUE string = "index-flights"

	FLIGTHS_DATA_ORIGIN string = "http://127.0.0.1:4004/"
	DB_URL              string = "http://elasticsearch:9200"

	DEFAULT string = "default"
)
