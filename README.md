# BackEndMLProject


Description:
This is a project which will detect the fire compliant building based on the number of smoke_detector, new_batteries, abc_extinguisher and clear_exit_routes


Tech Stack: 
Python FastAPI
Go
Postgresql
Docker
HTML for UI


Usage:
git clone the repo
docker compose up --build
Go to localhost:8080 to access the UI


Project structure:
BackEndMLProject
	|-fastapi_service
		|-db
		|-tests
	|-go_service
	    |-main.go
		|-tests
	|-docker-compose.yml
	|-README.md

	
Configuration:
Create the .env file in same location as docker-compose with the following values
	POSTGRES_USER=<your postgres user>
	POSTGRES_PASSWORD=<your postgres password>
	POSTGRES_DB=<your postgres DB>
	
