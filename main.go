package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"math"
)

type Position struct {
	X float64
	Y float64
	PathId int64
	Timestamp int64
}


func getCurrentPosition(path []Position, tNow int64) (Position, bool){
	//log.Println("         >> ----------- tNow:", tNow)
	if len(path) > 1{
		a := path[0]
		for _, b := range path[1:] {
			//log.Printf("         >> ---------- Pos: %+v / %+v", a, b)
			if a.Timestamp <= tNow && b.Timestamp > tNow{
				part := float64(tNow - a.Timestamp) / float64(b.Timestamp - a.Timestamp)
				//log.Println("             >> ------ part:", part)
				x := a.X + part * (b.X - a.X)
				y := a.Y + part * (b.Y - a.Y)
				//log.Println("             >> ------ x / y:", x, y)
				p := Position{X: x, Y: y, Timestamp: tNow, PathId: a.PathId}
				return p, true
			}
			a = b
		}
	}
	return Position{}, false
}

func main(){
	tStep := int64(1)
	minDist := float64(100)
	allPositions := make(map[int64][]Position)
	allPathes := make(map[int64]bool)

	file, err := os.Open("test_2.list")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	tMin := int64(0)
	tMax := int64(-1)
	for scanner.Scan() {
		ln := scanner.Text()
		data := strings.Split(ln, " ")
		pathId, _ := strconv.ParseInt(data[0], 10, 0)
		positionX, _ := strconv.ParseFloat(data[1], 0)
		positionY, _ := strconv.ParseFloat(data[2], 0)
		timestamp, _ := strconv.ParseInt(data[3], 10, 0)

		_, ok := allPathes[pathId]
		if !ok {
			allPathes[pathId] = true
			allPositions[pathId] = make([]Position, 0)
		}

		p := Position{X: positionX, Y: positionY, Timestamp: timestamp, PathId: pathId}
		allPositions[pathId] = append(allPositions[pathId], p)

		if tMax < 0{
			tMax = timestamp
		} else {
			if tMax < timestamp {
				tMax = timestamp
			}
		}
		if tMin > timestamp {
			tMin = timestamp
		}
	}
	file.Close()

	for tNow := tMin; tNow < tMax; tNow += tStep {
		//log.Println("-------------------- Time:", tNow)
		tempPositions := make([]Position, 0)
		for pathId := range allPathes{
			//log.Println("   >> -------------- PathId:", pathId)
			pos, ok := getCurrentPosition(allPositions[pathId], tNow)
			if ok{
				tempPositions = append(tempPositions, pos)
				//log.Printf("     >> -------------- Pos: %+v", pos)
			}
		}

		sort.SliceStable(tempPositions, func(i, j int) bool {
			return tempPositions[i].X < tempPositions[j].X
		})

		//log.Println(tempPositions)
		if len(tempPositions) > 1{
			a := tempPositions[0]
			for _, b := range tempPositions[1:] {
				if math.Abs(b.X-a.X) < minDist {
					if math.Abs(b.Y-a.Y) < minDist {
						//log.Println("In time:", tNow, " 2 pathes in same pos:", a.PathId, b.PathId)
					}
				}
				a = b
			}
		}
	}
}
