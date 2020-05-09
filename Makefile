CC=g++
CFLAG=-Wall
BIN=bin
SOURCE=src
LIST=bubblesort heapsort insertionsort mergesort quicksort selectionsort

all: $(LIST)

%: $(SOURCE)/%.c*
	mkdir -p $(BIN)
	$(CC) $< $(CFLAGS) -o $(BIN)/$@

clean:
	rm -rf $(BIN)
