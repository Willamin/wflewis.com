.PHONY: build
build:
	bundle exec middleman build

.PHONY: serve
serve:
	middleman serve

.PHONY: deploy
deploy: build
	scp -r build/* root@wflewis.com:/usr/share/caddy/
