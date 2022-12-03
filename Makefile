build: bwin blinux

bwin:
	@echo Building App binary for windows...
	set GOOS=windows&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o gvm.exe .
	@echo Done!

blinux:
	@echo Building App binary for linux...
	set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o gvm-linux .
	@echo Done!

setbin: build
	sudo cp gvm-linux /bin
	sudo mv /bin/gvm-linux /bin/gvm

test_switch: build
	gvm switch latest

test_dl: build
	gvm manager dl 1.18.5
	
test_use: build
	gvm manager use 1.18.5

test_scan: build
	gvm manager scan