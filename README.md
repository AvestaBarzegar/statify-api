# statify-api
The WIP API that will power Statify, the iOS application that gives you insight into your music listening history

use `docker compose up --build` to run the api locally

main.go is the entry point of the app

go to the server folder to see how I mount endpoints (i basically declare the URLs that the iOS app and possible website would call)

inside controllers i put the logic for handling http requests<br />
helpers are just helpers that help stuff (functions that make http requests, constants, etc)<br />
migrations is DB related but basically i have a migration to setup my databases and another to reverse the database setup<br />
models is where structs that represent information is stroed

