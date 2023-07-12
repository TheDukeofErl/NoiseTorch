NAME_SUFFIX=
WEBSITE_URL=https://github.com/noisetorch/NoiseTorch

VERSION := $(shell git describe --tags)

dev: rnnoise
	mkdir -p bin/
	go generate
	go build -ldflags '-X main.nameSuffix=${NAME_SUFFIX}_(dev) -X main.version=${VERSION} -X main.websiteURL=${WEBSITE_URL}' -o bin/noisetorch
release: rnnoise
	mkdir -p bin/
	mkdir -p tmp/

	mkdir -p tmp/.local/share/icons/hicolor/256x256/apps/
	cp assets/icon/noisetorch.png tmp/.local/share/icons/hicolor/256x256/apps/

	mkdir -p tmp/.local/share/applications/
	cp assets/noisetorch.desktop tmp/.local/share/applications/

	mkdir -p tmp/.local/bin/
	go generate
	CGO_ENABLED=0 GOOS=linux go build -trimpath -tags release -a -ldflags '-s -w -extldflags "-static" -X main.nameSuffix=${NAME_SUFFIX} -X main.version=${VERSION} -X main.distribution=official -X main.websiteURL=${WEBSITE_URL}' .
	mv noisetorch tmp/.local/bin/
	cd tmp/; \
	tar cvzf ../bin/NoiseTorch_x64_${VERSION}.tgz .
	rm -rf tmp/
rnnoise:
	git submodule update --init --recursive
	$(MAKE) -C c/ladspa
