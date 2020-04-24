PRODUCT=syslog_tester
VERSION=0.1

all:
	@echo "Options:"

dev:
	cd ../ && \
		docker build -f test_utils/Dockerfile --target dev --rm -t \
			$(PRODUCT)_dev:$(VERSION) .
	docker run --rm -it \
		--name $(PRODUCT)_dev \
		$(PRODUCT)_dev:$(VERSION) /bin/zsh

prod:
	cd ../ && \
		docker build -f test_utils/Dockerfile --target prod --rm -t \
		$(PRODUCT):$(VERSION) .

run_prod:
	docker run --rm -it \
		--name $(PRODUCT) \
		$(PRODUCT):$(VERSION)


attach_prod:
	docker exec -it \
		$(PRODUCT) /bin/zsh
