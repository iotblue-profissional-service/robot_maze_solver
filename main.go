package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Room struct {
	ID    string  `json:"id"`
	Left  *string `json:"left"`
	Right *string `json:"right"`
}

type Tree struct {
	Rooms map[string]*Room
	Path  []string
}

func NewTree(rooms []Room) *Tree {
	m := &Tree{
		Rooms: make(map[string]*Room),
		Path:  make([]string, 0),
	}
	for i := range rooms {
		m.Rooms[rooms[i].ID] = &rooms[i]
	}
	return m
}

func (m *Tree) findExit(current, prev string) bool {
	if current == "25" {
		m.Path = append(m.Path, "enters room "+current, "finish")
		return true
	}
	if current == "" || current == prev {
		return false
	}

	room, exists := m.Rooms[current]
	if !exists || room.Left == nil || room.Right == nil {
		return false
	}

	m.Path = append(m.Path, "enters room "+current)

	m.Path = append(m.Path, "move left")
	if m.findExit(*room.Left, current) {
		return true
	}
	m.Path = append(m.Path[:len(m.Path)-1])

	m.Path = append(m.Path, "move right")
	if m.findExit(*room.Right, current) {
		return true
	}
	m.Path = append(m.Path[:len(m.Path)-1])

	m.Path = append(m.Path, "go back")
	return false
}

func main() {
	router := gin.Default()

	router.POST("/findExit", func(c *gin.Context) {
		var rooms []Room
		if err := c.ShouldBindJSON(&rooms); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		maze := NewTree(rooms)
		maze.findExit("0", "")

		c.JSON(http.StatusOK, gin.H{
			"steps": maze.Path,
		})
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server Error")
		return
	}
}
