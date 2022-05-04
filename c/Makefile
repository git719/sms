TARGET  := smsclic
VERSION := 1.1.2

CC      = gcc
CFLAGS  = -g -Wall

.PHONY: release install clean

SRCS = $(wildcard *.c)
OBJS = $(SRCS:.c=.o)

LDFLAGS = -lcurl

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)

release:
	tar czf $(TARGET)-$(VERSION).tar.gz $(TARGET)
	shasum -a 256 $(TARGET)-$(VERSION).tar.gz

install:
	install $(TARGET) /usr/local/bin/$(TARGET)

clean:
	rm -f $(TARGET) $(TARGET)*.gz $(OBJS)
