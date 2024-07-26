package domain

type Config struct {
	AuthToken string `pulumi:"authToken"`
	Version *string `pulumi:"version,optional"`
} 
