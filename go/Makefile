# Makefile
# Assumes GOPATH is set up properly, e.g., $HOME/go

TARGET := smscli
VERSION := 1.1.0

default:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(TARGET)

all:
	rm -f $(TARGET) $(TARGET)*.gz
	go get -u github.com/vaughan0/go-ini
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(TARGET)

release:
	tar czf $(TARGET)-$(VERSION).tar.gz $(TARGET)
	shasum -a 256 $(TARGET)-$(VERSION).tar.gz

install:
	./install.sh $(TARGET)

clean:
	rm -f $(TARGET) $(TARGET)*.gz
