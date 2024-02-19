# Rates

A web app for loading, storing, converting and analyzing official currency 
rates from the National Bank of Moldova.<br>
The app uses Golang for backend, SolidJS for frontend and PostgreSQL for storage.

## Usage

Run the containers to start the database and http server.
Open localhost:8080 in your browser, load currency rates data.

## Available Scripts

In the main directory run:

### `make run`

Will build the application backend and frontend. Will start 2 containers:
- Postgres database
- Backend and frontend container

## Deployment

You can deploy the project using Docker compose. 
