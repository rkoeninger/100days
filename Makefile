bin:
ifeq ($(OS),Windows_NT)
	[ -d bin ] || mkdir bin
else
	mkdir -p bin
endif

hanoi: bin
	ozc -c ./src/hanoi.oz -o ./bin/hanoi.ozf
	ozengine ./bin/hanoi.ozf

mcm: bin
	ozc -c ./src/mcm.oz -o ./bin/mcm.ozf
	ozengine ./bin/mcm.ozf

permute: bin
	ozc -c ./src/permute.oz -o ./bin/permute.ozf
	ozengine ./bin/permute.ozf

clean:
	rm -rf bin
