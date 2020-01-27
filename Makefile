run:
	go run main.go

build:
	mkdir -p bin/
	go build -o bin/bmo

release:
	echo TODO

update:
	git pull
	go build
	cp proxima ~/Desktop/proxima
	pkill -f Desktop/proxima
	echo "Launch app manually"
