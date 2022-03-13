NAME:=burgerstore
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	go mod tidy
	go build -o bin/burgerstore .

run :
	./bin/burgerstore &

clean :
	kubectl delete -f mycrd >>/dev/null
	kubectl delete -f manifests >>/dev/null

test:
	go test -v ./...

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)

test-container :
	@docker rm -f $(NAME) || true
	@docker run \
	-v /Users/sdillikar/kubeconfig-kind-local-ip-cl1.yaml:/config/kubeconfig-kind-local-ip-cl1.yaml \
	--name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION) \
	/config/kubeconfig-kind-local-ip-cl1.yaml mydev-ns
	@echo  ""
	@sleep 2
	@docker stop $(NAME) || true
	@docker rm -f $(NAME) || true


