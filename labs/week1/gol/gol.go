package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	height := p.imageHeight
	width := p.imageWidth
	newWorld := make([][]byte, height, width)
	for i, a := range world{
		newA := make([]byte, width)
		for j, b := range a{
			newB := byte(0)
			numberLiving := getNeighbours(world, height, width, i, j)
			if b == 0 {
				if numberLiving == 3 {
					newB = 255
				} else{
					newB = 0
				}
			} else {
				if numberLiving < 2{
					newB = 0
				} else if numberLiving > 3{
					newB = 0
				} else{
					newB = 255
				}
			}
			newA[j] = newB
		}
		newWorld[i] = newA
	}
	return newWorld
}

func getNeighbours(world [][]byte, height int, width int, i int, j int) int{
	numberLiving := 0
	up := (i + 1) % height
	down := (i - 1) % height
	if down < 0 {
		down += height
	}
	right := (j + 1) % width
	left := (j -1) % width
	if left < 0 {
		left += width
	}
	if world[up][j] != 0{
		numberLiving++
	}
	if world[down][j] != 0{
		numberLiving++
	}
	if world[i][right] != 0{
		numberLiving++
	}
	if world[i][left] != 0{
		numberLiving++
	}
	if world[up][right] != 0{
		numberLiving++
	}
	if world[up][left] != 0{
		numberLiving++
	}
	if world[down][right] != 0{
		numberLiving++
	}
	if world[down][left] != 0{
		numberLiving++
	}
	return numberLiving
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	cells := []cell{}
	for i, a := range world{
		for j, b := range a{
			if b == 255{
				coordinates := cell{j, i}
				cells = append(cells, coordinates)
			}
		}
	}
	return cells
}
