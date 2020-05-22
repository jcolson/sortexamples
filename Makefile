CC=g++
GO=go
CFLAG=-Wall
BIN=bin
SOURCE=src
LIST=c-bubblesort c-heapsort c-insertionsort c-mergesort c-quicksort c-selectionsort go-heapsort

all: $(LIST)

c-%: $(SOURCE)/%.c*
	mkdir -p $(BIN)
	$(CC) $< $(CFLAGS) -o $(BIN)/$@

go-%: $(SOURCE)/%.go
	mkdir -p $(BIN)
	$(GO) build -o $(BIN)/$@ $<

clean:
	rm -rf $(BIN)
