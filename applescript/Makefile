TARGET := sms
VERSION := 1.1.0

default:
	ls

release:
	tar czf $(TARGET)-$(VERSION).tar.gz $(TARGET)
	shasum -a 256 $(TARGET)-$(VERSION).tar.gz

install:
	./install.sh $(TARGET)

clean:
	rm -f $(TARGET)-$(VERSION).tar.gz
