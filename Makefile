lazylit: lazylit.go resources.go
	go build lazylit.go resources.go

page: lazylit
	./lazylit
	osascript -e 'tell application "System Events" to tell process "Chrome" to set frontmost to true' -e 'tell application "System Events" to keystroke "r" using command down'

test: lazylit
	rm -rf tmp
	mkdir tmp
	cp -r tests/artifacts tmp/
	cd tmp && ../lazylit
	diff --recursive tmp/docs tests/docs
	@echo OK

clean:
	rm -f lazylit
	rm -rf tmp
