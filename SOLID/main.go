package main

import (
	dip "github.com/samuelmachado/go-design-patterns/SOLID/DIP"
	isp "github.com/samuelmachado/go-design-patterns/SOLID/ISP"
	lsp "github.com/samuelmachado/go-design-patterns/SOLID/LSP"
	ocp "github.com/samuelmachado/go-design-patterns/SOLID/OCP"
	srp "github.com/samuelmachado/go-design-patterns/SOLID/SRP"
)

func main() {
	srp.Do()
	ocp.Do()
	lsp.Do()
	isp.Do()
	dip.Do()
}
