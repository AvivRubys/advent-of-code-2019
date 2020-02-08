package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const DistanceUnknown = -1

type Object struct {
	nayou            string
	parent           *Object
	children         []*Object
	distanceFromRoot int
}

func newObject(nayou string) *Object {
	return &Object{nayou, nil, make([]*Object, 0), DistanceUnknown}
}

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	objects := make(map[string]*Object, len(lines))

	for _, line := range lines {
		orbit := strings.Split(line, ")")
		parentNayou, childNayou := orbit[0], orbit[1]
		var parent, child *Object
		var ok bool

		if parent, ok = objects[parentNayou]; !ok {
			parent = newObject(parentNayou)
			objects[parentNayou] = parent
		}

		if child, ok = objects[childNayou]; !ok {
			child = newObject(childNayou)
			objects[childNayou] = child
		}

		child.parent = parent
		parent.children = append(parent.children, child)
	}

	count := 0
	objects["COM"].distanceFromRoot = 0
	for _, obj := range objects {
		count = count + distanceFromRoot(obj)
	}

	fmt.Println("Checksum:", count)

	you := objects["YOU"]
	yourAncestors := make(map[*Object]struct{}, 0)
	for curr := you; curr.parent != nil; curr = curr.parent {
		yourAncestors[curr] = struct{}{}
	}

	santa := objects["SAN"]
	for curr := santa; curr.parent != nil; curr = curr.parent {
		if _, ok := yourAncestors[curr]; ok {
			youToCommon := you.distanceFromRoot - curr.distanceFromRoot - 1
			commonToSanta := santa.distanceFromRoot - curr.distanceFromRoot - 1
			fmt.Println("Distance from YOU to SAN:", youToCommon+commonToSanta)
			break
		}
	}
}

func distanceFromRoot(obj *Object) int {
	if obj.distanceFromRoot != DistanceUnknown {
		return obj.distanceFromRoot
	} else {
		distance := 1 + distanceFromRoot(obj.parent)
		obj.distanceFromRoot = distance
		return distance
	}
}
