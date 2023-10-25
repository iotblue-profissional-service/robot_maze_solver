package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Room struct {
	RoomId int
	Left   *Room
	Right  *Room
}

var visitedRooms = make(map[int]bool)

func iddfs(currentRoom *Room, goalRoomId int, path []string) []string {
	// Mark the current room as visited
	visitedRooms[currentRoom.RoomId] = true

	// Base case: if the current room is the goal room, return the path
	if currentRoom.RoomId == goalRoomId {
		return append(path, "finish")
	}

	// Try the left door
	if currentRoom.Left != nil && !visitedRooms[currentRoom.Left.RoomId] {
		path := iddfs(currentRoom.Left, goalRoomId, append(path, "move left", "enters room "+fmt.Sprint(currentRoom.Left.RoomId)))
		if path != nil {
			return path
		}
	}

	// Try the right door
	if currentRoom.Right != nil && !visitedRooms[currentRoom.Right.RoomId] {
		path := iddfs(currentRoom.Right, goalRoomId, append(path, "move right", "enters room "+fmt.Sprint(currentRoom.Right.RoomId)))
		if path != nil {
			return path
		}
	}

	// If IDDFS didn't find the exit room, return nil
	return nil
}

func handleRooms(c *gin.Context) {
	var rooms []Room
	if err := c.ShouldBindJSON(&rooms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a map to quickly access rooms by their ID
	roomMap := make(map[int]*Room)
	for i := range rooms {
		roomMap[rooms[i].RoomId] = &rooms[i]
	}

	// Link the rooms
	for i := range rooms {
		if rooms[i].Left != nil {
			rooms[i].Left = roomMap[rooms[i].Left.RoomId]
		}
		if rooms[i].Right != nil {
			rooms[i].Right = roomMap[rooms[i].Right.RoomId]
		}
	}

	// Run IDDFS for each unvisited room and keep track of the shortest path
	var shortestPath []string
	for i := range rooms {
		if !visitedRooms[rooms[i].RoomId] {
			path := iddfs(&rooms[i], 25, []string{"start from room " + fmt.Sprint(rooms[i].RoomId)})
			if path != nil && (shortestPath == nil || len(path) < len(shortestPath)) {
				shortestPath = path
			}
		}
	}

	if shortestPath == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No path to room 25"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"steps": shortestPath})
}

func main() {
	r := gin.Default()
	r.POST("/rooms", handleRooms)
	r.Run(":8080")
}
