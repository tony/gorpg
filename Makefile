all:
	go build

watch_test:
	while sleep 1; do ls *.go | entr -c go test; done
