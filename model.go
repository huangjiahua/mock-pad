package main

type Session struct {
	id        string
	owner     string
	problem   string
	instances []string
}

type Owner struct {
	name     string
	sessions []string
}
