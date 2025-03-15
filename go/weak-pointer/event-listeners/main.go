package main

import (
	"fmt"
	"weak"
)

type EventManager struct {
	listeners []weak.Pointer[func()]
}

func (em *EventManager) AddListener(listener func()) {
	em.listeners = append(em.listeners, weak.Make(&listener))
}

func (em *EventManager) TriggerEvent() {
	for _, weakListener := range em.listeners {
		if listener := weakListener.Value(); listener != nil {
			(*listener)()
		}
	}
}

func main() {
	em := &EventManager{}
	em.AddListener(func() { fmt.Println("Event received!") })

	em.TriggerEvent()

	em.listeners = nil
	fmt.Println("Clearing listeners...")

	em.TriggerEvent()
}
