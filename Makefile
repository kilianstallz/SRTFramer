commit:
	@cz commit

bump:
	@cz bump -at
	@git push --follow-tags origin main

test:
	@go test -race ./...
	make lint sec

sec:
	go run github.com/securego/gosec/v2/cmd/gosec@latest -quiet -severity medium ./...

lint:
	go run github.com/go-critic/go-critic/cmd/gocritic@latest check ./...

lint-hard:
	go run github.com/go-critic/go-critic/cmd/gocritic@latest check -enableAll ./...
sec-hard:
	go run github.com/securego/gosec/v2/cmd/gosec@latest -quiet ./...
