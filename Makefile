
install:
	cd cmd/sunstone-action && go install

# build base image
image:
	docker build -t sunstonedev/actions:0.1.0 -f Dockerfile.base .

push: image
	docker push sunstonedev/actions:0.1.0