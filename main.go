package main

import (
	"os"
	"fmt"
	"encoding/json"
	"errors"
	"sync"
)

type System struct {
	Uuid		string `json:"id"`
	Name		string `json:"name"`
	Region		string `json:"region"`
	State		string `json:"state"`
	AuxState	string `json:"auxState"`
	Endpoint	string `json:"apiEndpoint"`
	Edition		string `json:"systemType"`
	Version		string `json:"systemVersion"`
	// ConfigEndpoint string `json:"config_endpoint"`
}


var ErrorEmptySystemList = errors.New("systems list is empty")
var ErrorNameNotFound = errors.New("name not found")


func getData(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, err
}

func AllSystems(fileName string) ([]System, error) {
	data, err := getData(fileName)
	if err != nil {
		return nil, err
	}

	//Load json data into map to be able to extract from the items array
	var jsonData map[string][]System
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	//extract the items array
	systems := jsonData["items"]

	if systems == nil {
		return nil, ErrorEmptySystemList
	}

	return systems, nil
}

func searchName(name string, systems []System) *System {
	for _, s := range systems {
		if s.Name == name {
			return &s
		}
	}

	return nil
}

// Non Concurrent Function
// func SystemsByName(fileName string, names []string) (systems []System, err error) {
// 	allSystems, err := AllSystems(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for _, name := range names {
// 		if system := searchName(name, allSystems); system != nil {
// 			systems = append(systems, *system)
// 		} else {
// 			return systems, fmt.Errorf("%w: %s", ErrorNameNotFound, name)
// 		}
// 	}
//
// 	return systems, nil
// }



//Concurrent
//Comment the function abote and uncoment this one to test
func SystemsByName(fileName string, names []string) (systems []System, err error) {
	allSystems, err := AllSystems(fileName)
	if err != nil {
		return nil, err
	}

	systemsChan := make(chan System, len(names))
	errorChan := make(chan error, len(names))
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go func(systems []System, name string) {
			defer wg.Done()

			if system := searchName(name, allSystems); system != nil {
				systemsChan <- *system
			} else {
				errorChan <- fmt.Errorf("%w: %s", ErrorNameNotFound, name)
			}
		}(allSystems, name)
	}

	go func() {
		wg.Wait()
		close(systemsChan)
		close(errorChan)
	}()

	for system := range systemsChan {
		systems = append(systems, system)
	}

	if len(errorChan) > 0 {
		return systems, <-errorChan
	}

	return systems, nil
}

func main() {
	file := "data/systems_1000000.json"
	names := []string{"system500000", "system50", "system0", "system99", "foo"}
	systems, err := SystemsByName(file, names)
	fmt.Println(err)
	fmt.Println(systems)
}
