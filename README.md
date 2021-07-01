# DemoLogin

This project is generated to demonstrate the login functionality only.
It has below components
  - demo-login-ui - is responsible for UI
  - demo-login-server is for backend code which interacts with the db to register and login a user

## DemoLogin UI

Built using latest angular
Upon successful registration of a user redirects the user to /login page to login.
Upon successful login it stores a JWT token in cookie named "usertoken"
When User clicks on logout it delets the token from the browser cookie so that it can't be accessible any more.
When user is trying to logout then JWT Token needs to be passed to backend to authorize the user

## DemoLogin server

Built using go
Database used : MongoDB
Router used: gorilla mux
middleware used: negroni
Assumptions:
  Used rsa private and public key pair to sign and verify a user as a admin.
  Keys are stored in the code base(is not a good approach needs to be passed as env variables in production).
Built using docker

## How to run the application
As it is built using docker-compose so "docker-compose up --build" command can be used to build all the reqired images with network so that they can smoothly interact among themselves
1. Created a mongoDb docker image
2. Created a serer docker image in which go backend code is running
3. Create a ui docker image in which angular frontend code is running

To access the UI access the below url in browser
"http://localhost:4200"

It will load the UI click on register button to register user for first time
Once registered it will redirect to the /login page where user can login
Once logged in it will load /success page
user can click on logout buttom to logout
