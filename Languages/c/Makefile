CC = gcc
CFLAGS = -Wall -Wextra -std=c11 -g
LDFLAGS = -lpthread

# Auto-discover all .c files (excluding legacy/)
SRCS := $(shell find . -path ./legacy -prune -o -name '*.c' -print)
BINS := $(patsubst ./%.c, bin/%, $(SRCS))

.PHONY: all clean

all: $(BINS)

bin/%: ./%.c
	@mkdir -p $(dir $@)
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS)

clean:
	rm -rf bin/

# Run all binaries
.PHONY: test
test: all
	@for bin in $(BINS); do \
		echo "--- Running $$bin ---"; \
		./$$bin; \
		echo ""; \
	done
