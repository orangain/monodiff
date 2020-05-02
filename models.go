package monodiff

// ProjectDependency represents dependency to other project.
type ProjectDependency struct {
	ProjectName string
}

// FilesDependency represents dependency to files.
type FilesDependency struct {
	GlobPattern string
}

// Project represents unit of build.
type Project struct {
	Name                string
	FilesDependencies   []FilesDependency
	ProjectDependencies []ProjectDependency
}

// Spec represents set of projects.
type Spec struct {
	Projects []*Project
}
