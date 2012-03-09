LDFLAGS = -framework CoreServices -framework Cocoa

fswatch: main.o
	$(CC) $(LDFLAGS) $^ -o $@

clean:
	rm main.o fswatch
