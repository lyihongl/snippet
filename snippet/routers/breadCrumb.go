package controllers

import (
	//"net/http"

	//"github.com/gorilla/mux"
)

const (
	ROOT    = iota
	LEVEL_1 = iota
	LEVEL_2 = iota
)


type BreadCrumb struct {
	Title, Path string
	Level int
}

type BreadCrumbList struct {
	Crumbs []BreadCrumb
	Levels map[string]int
}

func (b *BreadCrumbList) Push(crumb *BreadCrumb) {
	b.Crumbs = append(b.Crumbs, *crumb)
}

func (b *BreadCrumbList) DefineLevels(levels map[string]int) {
	b.Levels = levels
}

func (b *BreadCrumbList) Update(action string) {
	listContainsRoot := false
	for i, e := range b.Crumbs{
		if(e.Level <= b.Levels[action]) {
			b.Crumbs[len(b.Crumbs) - 1], b.Crumbs[i] = b.Crumbs[i], b.Crumbs[len(b.Crumbs)-1]
			b.Crumbs = b.Crumbs[:len(b.Crumbs) - 1]
		}
		if(e.Level < b.Levels[action]) {
			listContainsRoot = true
		}
	}
	if(!listContainsRoot) {
	}
}
