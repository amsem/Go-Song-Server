# Music CRUD REST API
This is a simple CRUD REST API for music written in Go. It allows for basic operations such as creating, reading, updating, and deleting music entries. This API is intended as a starting point for more complex music-related projects, and can be easily extended to include additional features.

## Getting Started
To get started with this API, you will need to have Go ,Git and Postman installed on your machine. Once you have these tools installed, you can clone the repository and run the API.

```bash
git clone https://github.com/amsem/Go-Song-Server.git
```
```bash
cd Go-Song-Server
```
```bash
go run main.go
```

### The API now will run at port 8000

## Endpoints
The following endpoints are available for interacting with the API :
1. `GET /` : This endpoint returns a simple message to confirm that the API is running.
2. `GET /musics` : This endpoint returns a list of all music entries in the API.
3. `GET /musics/{id}` : This endpoint returns a specific music entry, based on the ID passed in the URL.
4. `POST /musics` : This endpoint allows you to create a new music entry. The entry must be in JSON format, and include a "title" and "artist" field.
5. `PUT /musics/{id}` : This endpoint allows you to update an existing music entry, based on the ID passed in the URL. The updated entry must be in JSON format.
6. `DELETE /musics/{id}` : This endpoint allows you to delete an existing music entry, based on the ID passed in the URL.

## Note 
This is just a first version and it does not have a database integration. This can be added in further versions to persist the data.

## Contribuation
If you'd like to contribute to this project, please feel free to submit a pull request. We welcome contributions of all kinds, including bug reports, feature requests, and code improvements.

