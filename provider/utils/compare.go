package utils

import "github.com/Genez-io/pulumi-genezio/provider/domain"

func CompareProjects(old, new domain.Project) bool {
	if old.Name != new.Name {
		return false
	}

	if old.Region != new.Region {
		return false
	}

	return true
}
