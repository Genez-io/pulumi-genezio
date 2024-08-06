package utils

import "github.com/Genez-io/pulumi-genezio/provider/domain"

func CompareProjects(old, new domain.Project) bool {
	if old.Name != new.Name {
		return false
	}

	if old.Region != new.Region {
		return false
	}

	// if old.CloudProvider == nil {
	// 	if new.CloudProvider != nil && *new.CloudProvider != "genezio-cloud" {
	// 		return false
	// 	}
	// } else {
	// 	if new.CloudProvider != nil {
	// 		if *old.CloudProvider != *new.CloudProvider {
	// 			return false
	// 		}
	// 	} else {
	// 		if *old.CloudProvider != "genezio-cloud" {
	// 			return false
	// 		}
	// 	}
	// }

	if old.CloudProvider != new.CloudProvider {
		return false
	}

	return true
}
