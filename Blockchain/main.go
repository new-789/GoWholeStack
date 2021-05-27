package main

func main() {
	bc := NewBlockChain("1LFEaXkHpDKvAu6WtfTqPwSs15EyAPg6nF")
	cli := Cli{bc}
	cli.Run()
}
