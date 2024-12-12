package main

import (
	"errors"
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
)

// RpcEventHandler handles the RPC call from Lua and creates a buffer with pendulum data.
func RpcEventHandler(v *nvim.Nvim, args map[string]interface{}) error {
	return nil
}

func main() {
	log.SetFlags(0)

	// Redirect stdout to stderr
	stdout := os.Stdout
	os.Stdout = os.Stderr

	// Connect to Neovim
	v, err := nvim.New(os.Stdin, stdout, stdout, log.Printf)
	if err != nil {
		log.Fatal(err)
	}

	// Register the "oolong-nvim" RPC handler, which receives Lua tables
	v.RegisterHandler("oolong-nvim", func(v *nvim.Nvim, args ...interface{}) error {
		// Expecting the first argument to be a map (Lua table)
		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		// Parse the first argument as a map
		argMap, ok := args[0].(map[string]interface{})
		if !ok {
			return errors.New("expected a map as the first argument")
		}

		// Call the actual handler with the parsed map
		return RpcEventHandler(v, argMap)
	})

	// Run the RPC message loop
	if err := v.Serve(); err != nil {
		log.Fatal(err)
	}
}
