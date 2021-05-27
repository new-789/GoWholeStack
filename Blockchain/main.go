package main

func main() {
	bc := NewBlockChain("14PxkwD8cTpzNAT1PYXRwK4qRNbkBVtgFP")
	cli := Cli{bc}
	cli.Run()
}
