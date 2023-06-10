# Rates

A web app for loading, storing, converting and analyzing official currency 
rates from the National Bank of Moldova.<br>
The app uses Golang for backend, SolidJS for frontend and PostgreSQL for storage.

## Usage

Run the containers to start the database and http server.
Open localhost:8080 in your browser, load currency rates data.

## Available Scripts

In the main directory run:

### `docker compose up --build`

Will build the application backend and frontend. Will start 2 containers:
- Postgres database
- Backend and frontend container

To run the frontend localy run:

### `cd frontend`
### `npm dev` or `npm start`

Runs the frontend in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br>

### `npm run build`

Builds the frontend for production to the `dist` folder.<br>

## Deployment

You can deploy the project using Docker compose. 
