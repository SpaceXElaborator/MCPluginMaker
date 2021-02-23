package pluginproject

import (
	"errors"
	"log"
	"os"
	"strings"

	"path/filepath"

	PluginCommands "SpaceXElaborator/PluginMaker/Command"
	pluginitems "SpaceXElaborator/PluginMaker/Item"
)

// Projects Holds a list of Current Projects
type Projects struct {
	Projects []*Project
}

// projectExists checks if the project with the same name (folder) exists on the system
func (cont *Projects) projectExists(name string) bool {
	for _, f := range cont.Projects {
		if strings.EqualFold(f.Name, name) {
			return true
		}
	}
	return false
}

// GetProject Returns a pointer to a Project given a name
func (cont *Projects) GetProject(name string) *Project {
	for _, f := range cont.Projects {
		if strings.EqualFold(f.Name, name) {
			return f
		}
	}
	return nil
}

// RemoveProject Will remove a Project given a name and delete everything inside of its folder and the folder itself
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
		err = os.RemoveAll(filepath.Join("./projects/"+name, fName))
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

// CreateNewProject Creates a new Project and will return an error if that project exists. It will also create the main folder and the subfolders needed for Maven
func (cont *Projects) CreateNewProject(name, author, groupid, artifactid, description string) error {
	if cont.projectExists(name) {
		return errors.New("Project Already Exists")
	}
	os.MkdirAll("projects/"+name+"/src/main/java/com/"+author+"/net", os.ModePerm)
	xml := PomXML{groupid, artifactid, description}
	cont.Projects = append(cont.Projects, &Project{name, author, xml, []*PluginCommands.Command{}, []*pluginitems.CustomItem{}})
	return nil
}
