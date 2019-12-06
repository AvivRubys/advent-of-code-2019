package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const DistanceUnknown = -1

type Object struct {
	name             string
	parent           *Object
	children         []*Object
	distanceFromRoot int
}

func newObject(name string) *Object {
	return &Object{name, nil, make([]*Object, 0), DistanceUnknown}
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
		parentName, childName := orbit[0], orbit[1]
		var parent, child *Object
		var ok bool

		if parent, ok = objects[parentName]; !ok {
			parent = newObject(parentName)
			objects[parentName] = parent
		}

		if child, ok = objects[childName]; !ok {
			child = newObject(childName)
			objects[childName] = child
		}

		child.parent = parent
		parent.children = append(parent.children, child)
	}

	count := 0
	objects["COM"].distanceFromRoot = 0
	for _, obj := range objects {
		count = count + distanceFromRoot(obj)
	}

	fmt.Println(count)
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
