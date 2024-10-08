# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2024-09-26

### 🚀 Features

- *(helpers)* Create helper function to get environment variable
- *(main)* Create simple program
- *(models)* Add sample Postgres model
- *(database)* Add helper functions to connect/disconnect to Postgres
- *(routes)* Add new script to setup sample routes
- *(main)* Establish connection to Postgres DB
- *(models)* Add "label" field to postgres_sample model
- *(database)* Add helper functions to connect/disconnect to MongoDB
- *(main)* Establish connection to MongoDB
- *(models)* Add sample MongoDB model
- *(jobs)* Add sample job for postgres
- *(jobs)* Add sample job for mongo
- *(jobs)* Implement job scheduler
- *(models)* Add sample models
- *(routes)* Add route to fetch transformed data from Mongo collection
- *(main)* Run jobs in data handler
- *(database)* Add logging and error handling

### 🐛 Bug Fixes

- *(jobs)* Run jobs with delay and interval

### 🚜 Refactor

- *(main)* Reformat logger output
- *(main)* Setup sample routes
- *(database)* Get host from environment variables
- *(database)* Make mongo client to be publicly accessible
- *(database)* Update log messages
- *(jobs)* Change function name
- *(main)* Change job's function name
- *(main)* Run jobs
- *(routes)* Change to log instead of fmt
- *(jobs)* Modify code to handle end of job execution

### 📚 Documentation

- Update README with project details

### 🎨 Styling

- *(helpers)* Start comments with uppercase letter
- *(database)* Add comment
- *(database)* Add and modify comments
- *(models)* Remove unwanted comment
- *(jobs)* Add, remove and modify comments
- *(jobs)* Modify and remove comments

### ⚙️ Miscellaneous Tasks

- *(database)* Remove unwanted code
- *(main)* Remove unwanted code
- Create generate-changelog.yml

### Build

- Get required packages
- Dockerize application
- Add Postgres service
- Add ORM library
- Add Postgres driver
- Tidy up dependencies
- Make web service to be dependent on postgres-db service
- Add Mongo driver
- Add service for MongoDB
- Make web service to be dependent on mongo-db service
- Add cron package
- Tidy up dependencies

<!-- generated by git-cliff -->
