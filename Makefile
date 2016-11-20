REPONAME = commons
DOCKERIMAGENAME = benchflow/$(REPONAME)
VERSION = dev

.PHONY: all build_release_go

all: build_release_go

build_go:
	$(MAKE) -C ./docker/go
	$(MAKE) -C ./kafka/go
	$(MAKE) -C ./keyname-hash-generator/go
	$(MAKE) -C ./minio/go

build_release_go: build_go