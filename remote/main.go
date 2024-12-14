package main

import (
	"errors"
	"log"
	"oolong-nvim/internal"
	"os"

	"github.com/neovim/go-client/nvim"
)

// RpcEventHandler handles the RPC call from Lua and creates a buffer with pendulum data.
func RpcEventHandler(v *nvim.Nvim, args map[string]interface{}) error {
	return nil
}

func main() {
	// l, _ := os.Create("oolong-nvim.log")
	// log.SetOutput(l)

	// Redirect standard output to standard error
	stdout := os.Stdout
	os.Stdout = os.Stderr

	// Connect to Neovim
	v, err := nvim.New(os.Stdin, stdout, stdout, log.Printf)
	if err != nil {
		log.Fatal(err)
	}

	// Register the "oolong-search" RPC handler
	v.RegisterHandler("oolong-search", func(v *nvim.Nvim, args ...interface{}) error {
		// Expecting the first argument to be a map (Lua table)
		if len(args) < 1 {
			err := errors.New("not enough arguments")
			log.Println(err)
			return err
		}

		// Parse the first argument as a map
		// CHANGE: make current arg part of table
		parsedArgs, ok := args[0].(string)
		if !ok {
			err := errors.New("expected a map as the first argument")
			log.Println(err)
			return err
		}

		if err := internal.SearchHandler(v, parsedArgs); err != nil {
			log.Println(err)
			return err
		}

		// Call the actual handler with the parsed map
		// return RpcEventHandler(v, argMap)
		return nil
	})

	// Register commands

	// Run the RPC message loop
	if err := v.Serve(); err != nil {
		log.Fatal(err)
	}
}
