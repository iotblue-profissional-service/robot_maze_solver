# Robot Maze Solver
<br>
![img.png](img.png)
<br>
A robot enters a maze consists of rooms, each room has index, started from room 0 it is required to 
find the room 25 which has the exit door, rooms are not sorted, and
each room has two doors (left & right) each door leads either to 
another room or leads to nothing (blocked door).

In the main file there is a model for room struct contains right door 
& left doors pointers, it is required to make an api endpoint which will 
receive an array off rooms like this: <br>  
[{"id": "0", "left":"5", "right": ""}, {"id": "4", "left":"1", "right": "8"}]
<br>with 25 objects in the array return the path that the robot follows 
like this: <br>  
{ "steps" : ["start","move left", "enters room 5", "move right", 
"enters room 2", "go back", "enters room 5", "move left",
"enters room 25", "finish"]} <br> 


note : take care not to fall in an infinite loop because some path  
will make a circle like this : 6 left -> 10 right -> 3 left -> 6 again




