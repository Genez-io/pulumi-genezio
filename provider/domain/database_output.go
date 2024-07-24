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

type GetDatabasesResponse struct{
	Status string `json:"status"`
	Databases []DatabaseDetails `json:"databases"`
}

type DatabaseDetails struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Region string `json:"region"`
	Type string `json:"type"`
	ConnectionUrl *string `json:"connectionUrl,omitempty"`
}