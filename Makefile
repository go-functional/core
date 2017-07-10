.PHONY: test
test:
	go test $$(glide nv)

.PHONY: simplefunctor
simplefunctor:
	go build -o bin/simplefunctor ./examples/simplefunctor
	bin/simplefunctor
.PHONY: bigfunctor
bigfunctor:
	go build -o bin/bigfunctor ./examples/bigfunctor
	bin/bigfunctor
