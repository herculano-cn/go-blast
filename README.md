# Go Blast Framework 

Just a web framework, to have a blast!
High performance, less latency, minimalist Go web framework

## Featured command line

Create first Go Blast! project

`blast new <project>`

Then you need to rename .env-example file to .env. So you can update config/environments/development.yaml with your local database and server configuration

Create project PostgreSQL database. First, make sure you have postgreSQL driver on your environment. Then you can go bblast with

`blast db:create` or you can create your own database name `blast db:create --dbname <db_name>`

Run project 

`blast init dev`

### Generators feature

You can make your own model/repo/handler instructions, using:

`blast gen.json Book title:string description:string num_pages:int` 

Following inserting route configurations. You can access Eg: `localhost:8080/books`