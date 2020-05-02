package monodiff

import (
	"encoding/json"
	"io/ioutil"
)

// ProjectJSON represents a project configuration in monodiff.json
type ProjectJSON struct {
	Deps []string `json:"deps"`
}

func loadSpec() (*Spec, error) {
	return loadSpecFile("monodiff.json")
}

func loadSpecFile(fileName string) (*Spec, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	specJSON := map[string]ProjectJSON{}
	if err := json.Unmarshal(bytes, &specJSON); err != nil {
		return nil, err
	}

	projectNameMap := map[string]bool{}
	for projectName := range specJSON {
		projectNameMap[projectName] = true
	}
	projects := []*Project{}

	for projectName, projectJSON := range specJSON {
		projectDependencies := []ProjectDependency{}
		filesDependencies := []FilesDependency{{GlobPattern: projectName}} // project depends on its directory

		for _, dep := range projectJSON.Deps {
			if _, present := projectNameMap[dep]; present {
				projectDependencies = append(projectDependencies, ProjectDependency{
					ProjectName: dep,
				})
			} else {
				filesDependencies = append(filesDependencies, FilesDependency{
					GlobPattern: dep,
				})
			}
		}
		projects = append(projects, &Project{
			Name:                projectName,
			ProjectDependencies: projectDependencies,
			FilesDependencies:   filesDependencies,
		})
	}

	return &Spec{Projects: projects}, nil
}
