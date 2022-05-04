TARGET := sms.sh
VERSION := 1.0.0

default:
	ls

release:
	tar czf $(TARGET)-$(VERSION).tar.gz $(TARGET)
	shasum -a 256 $(TARGET)-$(VERSION).tar.gz

install:
	./install.sh $(TARGET)

clean:
	rm -f $(TARGET)-$(VERSION).tar.gz
