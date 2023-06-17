IMAGE_NAME ?= sblast
IMAGE_PORT ?= 8080

.PHONY: image
image:
	@docker build . -t $(IMAGE_NAME)

run: image
	@docker run -p 127.0.0.1:$(IMAGE_PORT):$(IMAGE_PORT) --log-driver=none -a stdin -a stdout -a stderr $(IMAGE_NAME) 

deploy:
	@fly deploy
