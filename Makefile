build:
	mkdir -p bin/
	go build -o bin/bmo

update:
	git pull
	go build
	cp proxima ~/Desktop/proxima
	pkill -f Desktop/proxima
	echo "Launch app manually"
