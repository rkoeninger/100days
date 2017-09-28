run:
	ozc -c hanoi.oz -o hanoi.oza
	ozengine ./hanoi.oza

clean:
	rm -rf *.ozf *.oza
