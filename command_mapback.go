package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

func commandMapBack(c *config) error {
	if len(c.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(c.previous)

	if err != nil {
		return fmt.Errorf("Unable to retrieve locations. Error: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Unable to retrieve locations. Error: %v", err)
	}

	if err != nil {
		return fmt.Errorf("Unable to read the body. Error: %v", err)
	}

	var mapData jsonMap
	err = json.Unmarshal(body, &mapData)
		if err != nil {
		return fmt.Errorf("Unable to read the body. Error: %v", err)
	}

	for _, place := range mapData.Results {
		fmt.Println(place.Name)
	}

	c.next = mapData.Next
	c.previous = mapData.Previous

	return nil

}