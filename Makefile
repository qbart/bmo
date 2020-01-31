run:
	go run main.go

build:
	mkdir -p bin/
	go build -o bin/bmo

update:
	git pull
	go build -o bin/bmo
	pkill -f Desktop/bmo
	cp bin/bmo ~/Desktop/bmo
	echo "Launch app manually"
