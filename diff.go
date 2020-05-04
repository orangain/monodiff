package monodiff

import (
	"errors"
	"sort"
)

func detectChanges(spec *Spec, changedFiles []string) ([]*Project, error) {
	projectOrderMap := map[*Project]int{}
	for i, project := range spec.Projects {
		projectOrderMap[project] = i
	}

	sortedProjects, err := sortProjects(spec.Projects)
	if err != nil {
		return nil, err
	}

	changedProjects := []*Project{}
	changedProjectNameMap := map[string]bool{}

	for _, project := range sortedProjects {
		isChanged := false

		for _, projectDependency := range project.ProjectDependencies {
			if _, present := changedProjectNameMap[projectDependency.ProjectName]; present {
				isChanged = true
				break
			}
		}

		if !isChanged {
			for _, filesDependency := range project.FilesDependencies {
				if matchPattern(filesDependency.GlobPattern, changedFiles) {
					isChanged = true
					break
				}
			}
		}

		if isChanged {
			changedProjects = append(changedProjects, project)
			changedProjectNameMap[project.Name] = true
		}
	}

	// Sort changedProjects to preserve original order
	sort.Slice(changedProjects, func(i, j int) bool {
		return projectOrderMap[changedProjects[i]] < projectOrderMap[changedProjects[j]]
	})
	return changedProjects, nil
}

// sortProjects sort project to ensure that a project appears after all of its dependencies using topological sort.
// See: https://en.wikipedia.org/wiki/Topological_sorting#Depth-first_search
func sortProjects(projects []*Project) ([]*Project, error) {
	projectMap := map[string]*Project{}
	for _, project := range projects {
		projectMap[project.Name] = project
	}

	sortedProjects := []*Project{}
	processed := make(map[*Project]bool, len(projects)) // permanent mark
	seen := make(map[*Project]bool, len(projects))      // temporary mark

	var visit func(project *Project) error
	visit = func(project *Project) error {
		if _, present := processed[project]; present {
			return nil
		}
		if _, present := seen[project]; present {
			return errors.New("circular dependency detected")
		}

		seen[project] = true

		for _, dep := range project.ProjectDependencies {
			err := visit(projectMap[dep.ProjectName])
			if err != nil {
				return err
			}
		}

		delete(seen, project)
		processed[project] = true
		// Append project at last to ensure that a project appears after all of its dependencies.
		sortedProjects = append(sortedProjects, project)
		return nil
	}

	for _, project := range projects {
		if _, present := processed[project]; !present {
			err := visit(project)
			if err != nil {
				return nil, err
			}
		}
	}
	return sortedProjects, nil
}

func projectNames(projects []*Project) []string {
	names := make([]string, len(projects))
	for i, project := range projects {
		names[i] = project.Name
	}
	return names
}
