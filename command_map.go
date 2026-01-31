package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {

	//ask client for current/next page
	resp, err := cfg.ApiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	//print each location
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	//update page config
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil

}

func commandMapb(cfg *config, args ...string) error {
	//check if we're on first page or not
	if cfg.Previous == nil {
		fmt.Println("You're on first page")
		return nil
	}

	//ask client for current/previous page
	resp, err := cfg.ApiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	//print each location
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	//update page config
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil

}
