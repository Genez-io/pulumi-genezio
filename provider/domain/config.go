package domain

type Config struct {
	AuthToken string  `pulumi:"authToken"`
	Stage     *string `pulumi:"stage,optional"`
}
