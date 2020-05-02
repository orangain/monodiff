package monodiff

import (
	"reflect"
	"testing"
)

var directoryDependencySpec = &Spec{
	Projects: []*Project{
		{
			Name: "main",
			FilesDependencies: []FilesDependency{
				{"main"},
			},
		},
	},
}

var filesDependencySpec = &Spec{
	Projects: []*Project{
		{
			Name: "main",
			FilesDependencies: []FilesDependency{
				{"main"},
				{"package-lock.json"},
			},
		},
	},
}

var projectDependencySpec = &Spec{
	Projects: []*Project{
		{
			Name: "main",
			FilesDependencies: []FilesDependency{
				{"main"},
				{"package-lock.json"},
			},
			ProjectDependencies: []ProjectDependency{
				{"lib"},
			},
		},
		{
			Name: "lib",
			FilesDependencies: []FilesDependency{
				{"lib"},
				{"package-lock.json"},
			},
		},
	},
}

var complexProjectDependencySpec = &Spec{
	Projects: []*Project{
		{
			Name: "main",
			FilesDependencies: []FilesDependency{
				{"main"},
				{"package-lock.json"},
			},
			ProjectDependencies: []ProjectDependency{
				{"api"},
			},
		},
		{
			Name: "api",
			FilesDependencies: []FilesDependency{
				{"api"},
				{"package-lock.json"},
			},
			ProjectDependencies: []ProjectDependency{
				{"lib"},
			},
		},
		{
			Name: "lib",
			FilesDependencies: []FilesDependency{
				{"lib"},
				{"package-lock.json"},
			},
		},
	},
}

func TestDetectNoChangesWhenNoFilesAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(directoryDependencySpec, []string{})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(changedProjects) != 0 {
		t.Fatal("changedProjects must be empty")
	}
}

func TestDetectChangesWhenFilesInDirectoryAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(directoryDependencySpec, []string{"main/index.js"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if !reflect.DeepEqual(projectNames(changedProjects), []string{"main"}) {
		t.Fatal("changedProjects does not match")
	}
}

func TestDetectNoChangesWhenFilesOutOfDirectoryAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(directoryDependencySpec, []string{"lib/index.js"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(changedProjects) != 0 {
		t.Fatal("changedProjects must be empty")
	}
}

func TestDetectChangesWhenFilesInFilesDependencyAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(filesDependencySpec, []string{"package-lock.json"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if !reflect.DeepEqual(projectNames(changedProjects), []string{"main"}) {
		t.Fatal("changedProjects does not match")
	}
}

func TestDetectNoChangesWhenFilesOutOfFilesDependencyAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(filesDependencySpec, []string{"README.md"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(changedProjects) != 0 {
		t.Fatal("changedProjects must be empty")
	}
}

func TestDetectChangesWhenFilesInProjectDependencyAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(projectDependencySpec, []string{"lib/mail.js"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if !reflect.DeepEqual(projectNames(changedProjects), []string{"main", "lib"}) {
		t.Fatal("changedProjects does not match")
	}
}

func TestDetectNoChangesWhenFilesOutOfProjectDependencyAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(projectDependencySpec, []string{"README.md"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(changedProjects) != 0 {
		t.Fatal("changedProjects must be empty")
	}
}

func TestDetectChangesWhenFilesInComplexProjectDependencyAreChanged(t *testing.T) {
	changedProjects, err := detectChanges(complexProjectDependencySpec, []string{"lib/mail.js"})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	expected := []string{"main", "api", "lib"}
	actual := projectNames(changedProjects)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("changedProjects does not match.\nexpected: %#v\nbut got:  %#v", expected, actual)
	}
}
