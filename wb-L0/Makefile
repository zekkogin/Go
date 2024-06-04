.PHONY: all start_nats start_stan stop remove clean

include local.env
export $(shell sed 's/=.*//' local.env)

all: nats-server stan-server

nats-server:
	@echo "Starting NATS server on port $(NATS_PORT)"
	nohup nats-server -p $(NATS_PORT) -D > nats.log 2>&1 &

stan-server:
	@echo "Starting nats-streaming-server -cluster_id $(CLUSTER_ID) -store file -dir datastore -nats_server $(NATS_URL)"
	nohup nats-streaming-server -cluster_id $(CLUSTER_ID) -store file -dir datastore -nats_server $(NATS_URL) > stan.log 2>&1 &

clean: stop rm

stop:
	@pkill -f nats-server
	@pkill -f nats-streaming-server

rm:
	rm nats.log && rm stan.log && rm -rf datastore

pub:
	go run cmd/pub/main.go

sub:
	go run cmd/app/main.go

ps:
	ps aux | grep nats
