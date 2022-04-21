.PHONY:
.SILENT:



runStan:
		nats-streaming-server

runApp:
		go run ./cmd/web