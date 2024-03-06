IMAGE ?= go-webrtc-tunnel
VERSION ?= 0.0.1
ARCH  ?= $(shell $(GO) env GOARCH)

build-cli:
	mkdir -p bin
	go build -o bin/cli ./cmd/cli/main.go

build-signaling:
	mkdir -p bin
	go build -o bin/signaling ./cmd/signaling/main.go

test:
	@ if go test -v -count=1 ./pkg/...; then \
   		echo "\n\nTests passed"; \
   	else \
   		echo "\n\n!!! Tests failed !!!"; \
   		exit 1; \
   	fi

gen-grpc:
	docker run \
		-v .:/app \
		--workdir /app \
		--rm \
		bufbuild/buf generate proto