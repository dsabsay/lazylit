lazylit: lazylit.go resources.go
	go build lazylit.go resources.go

page: lazylit
	./lazylit
	osascript -e 'tell application "System Events" to tell process "Chrome" to set frontmost to true' -e 'tell application "System Events" to keystroke "r" using command down'
