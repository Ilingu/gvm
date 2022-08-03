GVM_BINARY=gvm.exe

build:
	@echo Building App binary...
	set GOOS=windows&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${GVM_BINARY} .
	@echo Done!

test_switch: build
	gvm switch latest

test_dl: build
	gvm manager dl 1.18.5
	
test_use: build
	gvm manager use 1.18.5

test_scan: build
	gvm manager scan