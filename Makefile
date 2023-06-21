PHOBY:run
run:
	go build && ./go-wire-demo

PHOBY:wire	
wire:
	wire gen ./...
