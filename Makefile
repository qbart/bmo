build:
	go build

update:
	git pull
	go build
	cp proxima ~/Desktop/proxima
	pkill -f Desktop/proxima
	echo "Launch app manually"
