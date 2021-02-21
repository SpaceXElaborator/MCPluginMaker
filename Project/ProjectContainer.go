package PluginProject

import (
	"strings"
	"errors"
	"os"
	"log"
	
	"path/filepath"
	
	"SpaceXElaborator/PluginMaker/Command"
	"SpaceXElaborator/PluginMaker/Item"
)

type Projects struct {
	Projects []*Project
}

func (cont *Projects) projectExists(name string) bool {
	for _, f := range cont.Projects {
		if strings.EqualFold(f.Name, name) {
			return true
		}
	}
	return false
}

func (cont *Projects) GetProject(name string) *Project {
	for _, f := range cont.Projects {
		if strings.EqualFold(f.Name, name) {
			return f
		}
	}
	return nil
}

func (cont *Projects) RemoveProject(name string) {
	var i int
	for ind, proj := range cont.Projects {
		if strings.EqualFold(proj.Name, name) {
			i = ind
		}
	}
	
	d, err := os.Open("./projects/" + name)
	if err != nil {
		log.Fatal(err)
	}
	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, fName := range names {
		err = os.RemoveAll(filepath.Join("./projects/" + name, fName))
		if err != nil {
			log.Fatal(err)
		}
	}
	d.Close()
	err = os.Remove("./projects/" + name)
	if err != nil {
		log.Fatal(err)
	}
	
	copy(cont.Projects[i:], cont.Projects[i+1:])
	cont.Projects = cont.Projects[:len(cont.Projects)-1]
}

func (cont *Projects) CreateNewProject(name, author, groupid, artifactid, description string) error {
	if cont.projectExists(name) {
		return errors.New("Project Already Exists")
	}
	os.MkdirAll("projects/" + name + "/src/main/java/com/" + author + "/net", os.ModePerm)
	xml := PomXML{groupid, artifactid, description}
	cont.Projects = append(cont.Projects, &Project{name, author, xml, []*PluginCommands.Command{}, []*PluginItems.CustomItem{}})
	return nil
}