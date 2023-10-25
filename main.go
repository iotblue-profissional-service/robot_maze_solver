package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Room struct {
	ID      string
	Left    string
	Right   string
	Visited bool //to know if this room has been visited or not
}

type Maze struct {
	Path  []string         // will write down the path of the solution , path : slice
	Rooms map[string]*Room //the key is string (IDs) and the value will be pointer to the Room struct
}

func NewMaze(data []byte) (*Maze, error) { // data : is the data that come from the json

	var rooms []Room                    // identify rooms as empty slice of the Room struct
	err := json.Unmarshal(data, &rooms) // convert the data that came from json and put it in this slice

	if err != nil {
		// if there is an error = err  & nil is the zero value of the slice
		return nil, err
	}

	maze := &Maze{ // create new object of Maze
		Rooms: make(map[string]*Room),
		Path:  make([]string, 0),
	}

	for i := range rooms {
		maze.Rooms[rooms[i].ID] = &rooms[i] // assign the IDs (type string ) into the new maze
		//we use a '&' in (&rooms[i]) to make the changes reflect to the original
	}

	return maze, nil //return the new maze
}

func (m *Maze) findExit(room *Room) bool {

	if room.ID == "25" {
		return true
	} // if the roomID wasn't 25 ( the exist room) --> mark it as visited room
	room.Visited = true

	// check if there is room in the left & search in the map if the room.left has been visited or not
	if room.Left != "" && !m.Rooms[room.Left].Visited {
		m.Path = append(m.Path, "Move Left", "Enters the room ", room.Left)
		// room.left is the type of string ,so it may be "5" for example
		// say to the robot to once it's in the current room to move left and enter room...

		if m.findExit(m.Rooms[room.Left]) { //recursive
			return true // that mean it finds left room , and it has the path to the exit
		}
		m.Path = m.Path[:len(m.Path)-2]
		// so the previous if didn't return true , then it didn't find the path to exist --> delete the last to steps in the path slice
	}

	//we will return the same if condition but on the right room
	if room.Right != "" && !m.Rooms[room.Right].Visited {
		m.Path = append(m.Path, "Move Right", "Enter the room ", room.Right)
		if m.findExit(m.Rooms[room.Right]) {
			return true
		}
		m.Path = m.Path[:len(m.Path)-2]
	}

	return false // false : still didn't find the exist room (25)
}

func (m *Maze) Solve() []string {
	if m.findExit(m.Rooms["0"]) {
		return append([]string{"Start "}, append(m.Path, "Finish")...) // to add the finish word to the end of the slice
		// add the start word to new slice , then add to it all the elements in the path slice
	}
	fmt.Println("there is no path to the exit")
	return []string{} // if no exit path found , return empty slice
}

func main() {
	reader := bufio.NewReader(os.Stdin) // create new reader with buffer
	fmt.Println("Enter Json data : ")   // enter it in one line
	data, _ := reader.ReadString('\n')  // to read the data

	maze, err := NewMaze([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	path := maze.Solve()
	fmt.Println(path)
}
