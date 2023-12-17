# Netflix API using GO and GOFR
This is simple a REST_API for managing the movies watched by you.It provides us with various endpoints to perform various operations such as retrieving all the movies or retreeiving a single movie by ID.It
can also be used for updating a movie using a particular ID or deleting a movie using a particular ID . It is implemented using GO and GOFR framework ,MONGODB and can perform CRUD operations.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://go.dev/doc/install)
- [MongoDB](https://www.mongodb.com/docs/manual/administration/install-community/)
  
## Getting Started

1. Clone the repository:

   ```bash
   https://github.com/Arshverma001/Simple-REST-API-using-go-and-gofr.git
   
   cd GOFR-Assignment
2. Download Dependencies:

   ```bash
   go mod download

3. Verify Dependencies:

    ```bash
   cat go.sum
4. Run Project
    
    ```bash
    go run main.go

5. Open Server
    
    ```bash
   http://localhost:8000/api/movies
    
## API Endpoints

The application provides RESTful API endpoints for CRUD operations on movie records.

### Base URL

The base URL for all endpoints is:

### List of Endpoints

#### 1. **Get All Movies**

- **Endpoint:**
  - `GET /api/movies`

- **Description:**
  - Retrieves a list of all movies.

- **Example Response:**
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/ee57ea10-0dde-47bd-81fc-c2432ccdb99b)


#### 2. **Get movie by ID**

- **Endpoint:**
  - `GET api/movie/:id`
- **Parameters:**
  - `id`: The unique identifier for each movie.

- **Description:**
  - Retrieves information about a specific movie based on the provided ID.

- **Example Response:**
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/204c23ce-bde9-4564-ae8b-3a826603a2dd)


#### 3. **Insert a NEW movie**

- **Endpoint:**
  - `POST /api/movie`
- **Request Body:**
  ```json
  [    { "movie":"How to train your dragon",  "watched":true }   ]  


- **Description:**
  - Add a new movie.
    
- **Example Response:**
  ```json
  {      "_id": "657f010d2a096dee359e244b",      "movie": "How to train your dragon",     "watched": true    }


![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/d184d7ee-bf8a-49a4-9ddf-ecef2f75d107)
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/65818db6-af29-4e73-90ca-0d5ef6c3cf12)



#### 4. **Update Movie by ID**

- **Endpoint:**
  - `PUT /api/movie/:id`
- **Parameters:**
  - `id`: The unique identifier for the movie.

- **Request Body:**
  ```json
  [   {"movie":"Cars","watched":true}   ]  

- **Description:**
  - Updates information about a specific movie based on the provided ID.
    
- **Example Response:**
  
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/2140545f-2097-40ee-8c89-1d0d7106434e)
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/a3f8810c-2122-40c5-96a3-a879cbdd4702)



 
#### 5. **Delete Movie by ID**

- **Endpoint:**
  - `DELETE /api/movie/:id`

- **Parameters:**
  - `id`: The unique identifier for the movie.

- **Description:**
  - Deletes a specific movie based on the provided ID.

- **Example Response:**

![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/41dafdd2-2755-46ef-a37b-7f0e3a9ab180)
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/b15bc87f-6201-45c5-938c-4696d056a3b5)

## UML Diagram
![image](https://github.com/Arshverma001/Simple-REST--API-using-go-and-gofr/assets/87807771/f1a590e4-c3fd-46cf-8743-55d2e0846872)


