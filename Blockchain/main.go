package main

func main() {
	bc := NewBlockChain()
	cli := Cli{bc}
	cli.Run()
	/*
	bc.AddBlock("111111111111111111111111111")
	bc.AddBlock("222222222222222222222222222")


	 */
}
