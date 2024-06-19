# Authenticated API project with OpenF1 
This project is to further my knowledge on GoLang and Postgresql by creating a backend that includes: Authorization, Third Part REST API connections, and Postgresql database configurations.

## Authorization
I used the bcrypt package and JWT (version 5) to create my auth middleware. I am taking the username and password from the request body, and hashing it into my postgresql users table. For logging in, I am comparing the passwords from the request body and the database to verify users and setting an Authorization cookie which has an expiration of a month. I am using the Gin Gonic framework to help make the API calls and handle request and response bodies. 

## OpenF1
I am using the OpenF1 API to get driver data for logged in users. An example curl for this API is: `curl "https://api.openf1.org/v1/car_data?driver_number=55&session_key=9159&speed>=315"`. I think it would be a cool v2 of this project to combine the driver, session, and car data endpoints to create custom api endpoints to be used by a frontend that includes dashboards on favorite driver or teams. 

## Postgresql
I setup my database using PGAdmin4. I am connecting to the database with the help of Gorm.
