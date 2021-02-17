package main

type Settings struct {
	Author string `json:"Author"`
	Dark bool `json:"Mode"`
}

type Command struct {
	Author, Name, SlashCommand string
}

type MainFile struct {
	Author string
	Cmds []Command
}

type PomXML struct {
	Author, Name, GroupID, ArtifactID, Description string
}

type Project struct {
	Name, Author string
	GroupID, ArtifactID, Description string
	Cmds []Command
}