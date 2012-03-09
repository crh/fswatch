LDFLAGS = -framework CoreServices -framework Cocoa

fswatch: fswatch.o main.o
	$(CC) $(LDFLAGS) $^ -o $@

testrunner: fswatch.o test.o
	$(CC) $(LDFLAGS) $^ -o $@

test: testrunner
	@./testrunner
	@printf "\e[1;32mPASS\e[0m\n"

clean:
	rm -f fswatch.o main.o test.o fswatch testrunner
