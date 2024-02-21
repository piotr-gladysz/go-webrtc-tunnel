IMAGE ?= go-webrtc-tunnel
VERSION ?= 0.0.1
ARCH  ?= $(shell $(GO) env GOARCH)

test:
	go test -v -count=1 ./pkg/...

gen-grpc:
	docker run \
		-v .:/app \
		--workdir /app \
		--rm \
		bufbuild/buf generate proto