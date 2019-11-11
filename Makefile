
install:
	cd cmd/sunstone-action && go install

# build base image
image:
	docker build -t sunstonedev/actions:latest -f Dockerfile.base .

push: image
	docker push sunstonedev/actions:latest