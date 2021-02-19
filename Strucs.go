package main

type Settings struct {
	Author string `json:"Author"`
	Dark bool `json:"Mode"`
}

type Command struct {
	Author, CommandType, Name, SlashCommand string
}

type CustomItem struct {
	Author, ItemMaterial, ItemName string
	ItemDescription []string
}

type ShapedCrafting struct {
	Author, Item string
	String1, String2, String3 string
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
	Items []CustomItem
}