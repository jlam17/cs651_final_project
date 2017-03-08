/*
	worker.go
	authors: Justin Chen, Johnson Lam
	
	Worker interface with party.

	3.7.17
*/

package main

import (
	"config"
	"fmt"
	"message"
)

func main () { 
	// Create config
	var conf = config.Configuration{}
	cfile   := "config.json"
	conf.Load(cfile)

	// Already joined a party
	if len(conf.Party.IP) > 0 && conf.Party.Port > 0 && len(conf.Party.Alias) > 0 {
		fmt.Println(conf.Party)
		message.JoinParty(&conf)
	} else {
		fmt.Println("specify party to join and complete config...")
	}
}