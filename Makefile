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

.PHONY: chanfunctor
chanfunctor: 
	go build -o bin/chanfunctor ./examples/chanfunctor
	bin/chanfunctor


.PHONY: optional
optional:
	go build -o bin/optional ./examples/optional
	bin/optional

.PHONY: either
either: 
	go build -o bin/either ./examples/either
	bin/either

.PHONE: typeclass-eq
typeclass-eq:
	go build -o bin/typeclass/eq ./examples/typeclass/eq
	bin/typeclass/eq
