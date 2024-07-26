package domain

type SubscriptionLimits struct {
	MaxProjects int `json:"maxProjects"`
	MaxInvocations int `json:"maxInvocations"`
	ExecutionTime float64 `json:"executionTime"`
	MaxConcurrency float64 `json:"maxConcurrency"`
	MaxCollaborators int `json:"maxCollaborators"`
}

type OnboardingInfo struct {
	OnboardingComplete bool `json:"onboardingComplete"`
	Role string `json:"role"`
	ProgrammingLanguages []string `json:"programmingLanguages"`
	ExperienceLevel string `json:"experienceLevel"`
}

type UserPayload struct {
	ID string `json:"id"`
	Email string `json:"email"`
	ProfileUrl string `json:"profileUrl"`
	SubscriptionPlan string `json:"subscriptionPlan"`
	SubscriptionPrice string `json:"subscriptionPrice"`
	MemberSince string `json:"memberSince"`
	SubscriptionLimits SubscriptionLimits `json:"subscriptionLimits"`
	CustomSubscription bool `json:"customSubscription"`
	OnboardingInfo OnboardingInfo `json:"onboardingInfo"`
}