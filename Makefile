run:
	go run main.go

build:
	mkdir -p bin/
	go build -o bin/bmo

update:
	git pull
	go build -o bin/bmo
	cp bin/bmo ~/Desktop/bmo
	pkill -f Desktop/bmo
	echo "Launch app manually"
