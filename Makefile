build_plugin:
	$(MAKE) -C plugin build

clean_plugin:
	$(MAKE) -C plugin clean

clean: clean_plugin
	-rm main && rm *.exe

build: clean build_plugin
	go build main.go