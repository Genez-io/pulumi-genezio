package domain

type CreateDatabaseResponse struct{
	DatabaseId string `json:"databaseId"`
	Name string `json:"name"`
	Region string `json:"region"`
	Type string `json:"type"`
}

type GetDatabaseConnectionUrlResponse struct{
	ConnectionUrl string `json:"connectionUrl"`
	Status string `json:"status"`
}