
NAME:=burgerstore
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	# go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
	go build -a -o bin/server_main ./*.go

clean :
	@rm bin/*

test:
	go test -v ./...  

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)	

test-container :
	@docker rm -f $(NAME) || true
	@docker run -dp 9001:9001 --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	curl -s localhost:9001	