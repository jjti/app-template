IMAGE_NAME ?= sblast

deploy: ui
	@fly deploy

.PHONY: ui
ui:
	cd ui && npm run build

image:
	@docker build . -t $(IMAGE_NAME)
