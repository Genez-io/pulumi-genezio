package utils

import (
	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func CompareProjects(old, new domain.Project) bool {
	if old.Name != new.Name {
		return false
	}

	if old.Region != new.Region {
		return false
	}

	return true
}

func CompareAuthProviders(old, new domain.AuthenticationProviders) bool {
	if old.Email == nil {
		if new.Email != nil && *new.Email {
			return false
		}
	} else {
		if new.Email != nil {
			if *old.Email != *new.Email {
				return false
			}
		} else {
			if *old.Email {
				return false
			}
		}
	}

	if old.Web3 == nil {
		if new.Web3 != nil && *new.Web3 {
			return false
		}
	} else {
		if new.Web3 != nil {
			if *old.Web3 != *new.Web3 {
				return false
			}
		} else {
			if *old.Web3 {
				return false
			}
		}
	}

	if old.Google == nil {
		if new.Google != nil {
			return false
		}
	} else {
		if new.Google != nil {
			if old.Google.ID != new.Google.ID || old.Google.Secret != new.Google.Secret {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
