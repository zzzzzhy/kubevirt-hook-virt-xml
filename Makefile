TAG ?= "latest"
REGISTRY ?= "docker.io/alicefr"
IMAGE ?= "virt-xml"
build: clean
	go build -o updater main.go

update:
	go mod tidy
	go mod vendor

container: build
	docker build -t $(REGISTRY)/$(IMAGE):$(TAG) .

push:
	docker push $(REGISTRY)/$(IMAGE):$(TAG)

clean: 
	rm -f updater
