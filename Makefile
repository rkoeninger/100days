bin-directory:
ifeq ($(OS),Windows_NT)
	[ -d bin ] || mkdir bin
else
	mkdir -p bin
endif

hanoi: bin-directory
	ozc -c ./src/hanoi.oz -o ./bin/hanoi.ozf
	ozengine ./bin/hanoi.ozf

clean:
	rm -rf bin
