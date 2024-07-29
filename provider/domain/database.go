package domain

type CreateDatabaseRequest struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Type   string `json:"type,omitempty"`
}

type CreateDatabaseResponse struct {
	DatabaseId string `json:"databaseId"`
	Status     string `json:"status"`
}

type GetDatabaseConnectionUrlResponse struct {
	ConnectionUrl string `json:"connectionUrl"`
	Status        string `json:"status"`
}

type GetDatabaseResponse struct {
	Status    string            `json:"status"`
	Databases []DatabaseDetails `json:"databases"`
}

type DatabaseDetails struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Region        string  `json:"region"`
	Type          string  `json:"type,omitempty"`
	ConnectionUrl *string `json:"connectionUrl,omitempty"`
}
