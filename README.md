![Go Blast Framework](data:image/svg+xml,<svg style="width:15em;" xmlns:mydata="http://www.w3.org/2000/svg" mydata:contrastcolor="ffffff" mydata:template="Contrast" mydata:presentation="2.5" mydata:layouttype="undefined" mydata:specialfontid="undefined" mydata:id1="640" mydata:id2="209" mydata:companyname="Go Blast!" mydata:companytagline="Do less, do better" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 45 550 380"><g fill="#e06287" fill-rule="none" stroke="none" stroke-width="1" stroke-linecap="butt" stroke-linejoin="miter" stroke-miterlimit="10" stroke-dasharray="" stroke-dashoffset="0" font-family="none" font-weight="none" font-size="none" text-anchor="none" style="mix-blend-mode: normal"><g data-paper-data="{&quot;isGlobalGroup&quot;:true,&quot;bounds&quot;:{&quot;x&quot;:64.99999999999999,&quot;y&quot;:181.21476832959485,&quot;width&quot;:420,&quot;height&quot;:107.5704633408103}}"><g data-paper-data="{&quot;isSecondaryText&quot;:true}" fill-rule="nonzero"><path d="M220.2955,263.48726h-2.9663c-0.38902,0 -0.52275,0.13373 -0.52275,0.51059v7.61025c0,0.37687 0.13373,0.51059 0.52275,0.51059h2.9663c2.39492,0 3.73218,-1.31295 3.73218,-3.40395v-1.82354c0,-2.10315 -1.33726,-3.40395 -3.73218,-3.40395zM221.62061,266.92768v1.7506c0,0.86314 -0.42549,1.31295 -1.41021,1.31295h-0.99687v-4.3765h0.99687c0.98471,0 1.41021,0.44981 1.41021,1.31295z" data-paper-data="{&quot;


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