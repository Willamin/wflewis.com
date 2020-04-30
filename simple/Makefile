.DEFAULT_GOAL := all

.PHONY: build_dir
build_dir:
	mkdir -p build

.PHONY: article
article: build_dir
	pandoc -f gfm -t html --metadata title="Will Lewis" --template="wfl-theme" > build/index.html

.PHONY: index
index: build_dir
	pandoc -t gfm -t html --metadata title="Will Lewis" --template="wfl-theme" > build/index.html < index.md

.PHONE: all
all: index
