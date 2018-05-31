package main

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("document").Call("getElementById", "reset").Set("onclick", restartGame)
	createBoard()
	refreshBoard()
}

func createBoard() {
	println("create board")
	usStates := []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "California",
		"Colorado", "Connecticut", "Delaware", "Florida", "Georgia",
		"Hawaii", "Idaho", "Illinois", "Indiana", "Iowa",
		"Kansas", "Kentucky", "Louisiana", "Maine", "Maryland",
		"Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri",
		"Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio",
		"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont",
		"Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming",
	}

	document := js.Global.Get("document")
	for _, name := range usStates {
		elem := document.Call("createElement", "div")
		elem.Set("className", "us_state")
		elem.Set("id", name)
		elem.Set("textContent", name)
		elem.Set("onclick", toggleFound)
		document.Call("getElementById", "content").Call("appendChild", elem)
	}
}

func toggleFound(e *js.Object) {
	toggleFoundness(e)
	refreshBoard()
}

func toggleFoundness(e *js.Object) {
	name := e.Get("srcElement").Get("id")
	item := js.Global.Get("localStorage").Call("getItem", name)
	if item == js.Undefined || item == nil {
		println("add", name, "to localStorage")
		js.Global.Get("localStorage").Call("setItem", name, true)
	} else {
		println("remove", name, "from localStorage")
		js.Global.Get("localStorage").Call("removeItem", name)
	}
}

func refreshBoard() {
	println("refresh board")
	nodeList := js.Global.Get("document").Call("querySelectorAll", ".us_state")
	length := nodeList.Get("length").Int()
	var count int
	for i := 0; i < length; i++ {
		elem := nodeList.Call("item", i)
		item := js.Global.Get("localStorage").Call("getItem", elem.Get("id"))
		stored := !(item == js.Undefined || item == nil)
		if stored {
			count++
		}
		found := elem.Get("classList").Call("contains", "found").Bool()
		switch {
		case stored && !found:
			println("add \"found\" class to", elem.Get("id"))
			elem.Get("classList").Call("add", "found")
		case !stored && found:
			println("remove \"found\" class from", elem.Get("id"))
			elem.Get("classList").Call("remove", "found")
		}
	}
	updateProgressBar(count)
}

func updateProgressBar(count int) {
	doneElem := js.Global.Get("document").Call("querySelector", ".progress-done")
	todoElem := js.Global.Get("document").Call("querySelector", ".progress-todo")

	if count >= 0 && count <= 50 {
		remaining := 50 - count
		doneElem.Get("style").Set("flex-grow", count)
		doneElem.Get("style").Set("width", count)
		doneElem.Set("textContent", strconv.Itoa(count))
		todoElem.Get("style").Set("flex-grow", remaining)
		todoElem.Get("style").Set("width", remaining)
		todoElem.Set("textContent", strconv.Itoa(remaining))
	}
	if count == 50 {
		doneElem.Set("textContent", "Winner!  You rock!")
	}
}

func restartGame() {
	println("unfinding all")
	js.Global.Get("localStorage").Call("clear")
	refreshBoard()
}
