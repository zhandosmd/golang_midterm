package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]string
}

type event struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var c = Container{
	counters: map[string]string{"a": "value a", "b": "value b"},
}

var wg sync.WaitGroup

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Midterm Project")
}

func (c *Container) updateValueReference(key string, newValue string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[key] = newValue
}

func (c *Container) deleteValueReference(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.counters, key)
}

var updateValue = func(key string, newValue string) {
	c.updateValueReference(key, newValue)
	wg.Done()
}

var deleteKey = func(key string) {
	c.deleteValueReference(key)
	wg.Done()
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	wg.Add(1)
	updateValue(newEvent.Key, newEvent.Value)
	wg.Wait()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for key, value := range c.counters {
		if key == eventID {
			var foundElement event
			foundElement.Key = key
			foundElement.Value = value
			json.NewEncoder(w).Encode(foundElement)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.counters)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for key := range c.counters {
		if key == eventID {
			wg.Add(1)
			updateValue(key, updatedEvent.Value)
			wg.Wait()
			var updatedElement event
			updatedElement.Key = key
			updatedElement.Value = updatedEvent.Value
			json.NewEncoder(w).Encode(updatedElement)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for key := range c.counters {
		if key == eventID {
			wg.Add(1)
			deleteKey(key)
			wg.Wait()
			fmt.Fprintf(w, "The event with KEY %v has been deleted successfully", eventID)
		}
	}
}
