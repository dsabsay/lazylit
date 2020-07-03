gocco: gocco.go resources.go
	go build gocco.go resources.go

page: gocco
	./gocco
	osascript -e 'tell application "System Events" to tell process "Chrome" to set frontmost to true' -e 'tell application "System Events" to keystroke "r" using command down'
