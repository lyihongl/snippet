package controllers

const(
	ROOT = iota
	LEVEL_1 = iota
	LEVEL_2 = iota
)

type BreadCrumb struct {
	Title, Path, Level string
}

type BreadCrumbList struct {
	Crumbs []BreadCrumb
}

func (b *BreadCrumbList) Push(crumb *BreadCrumb) {
	b.Crumbs = append(b.Crumbs, *crumb)
}
