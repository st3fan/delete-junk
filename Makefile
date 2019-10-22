NAME = delete-junk
TAG = st3fan/$(NAME):master

REGISTRY = registry.sateh.com

build:
	docker build -t $(TAG) .

push:
	docker tag $(TAG) $(REGISTRY)/$(TAG)
	docker push $(REGISTRY)/$(TAG)

run:
	docker run --rm -it --name $(NAME) $(TAG)

start:
	docker run -dit --name $(NAME) --restart unless-stopped $(TAG)
