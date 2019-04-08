# EZCTF-API

### Overview
---

This is the API written to accompany the ReactJS front end portion of the EZCTF - Webapp.

A continuation of [my EZCTF app from last year.](https://github.com/AnthonyLaiuppa/ezctf) 

I thought it would be fun to rewrite it with a different approach using Golang/ReactJS.

Since this is largely for learning somethings may not be quite perfect. 

In its current state authentication middleware is installed and basic CRUDs are present.

Next step is to add CTF related logic to the API.


##### Running
`export DBCONNSTRING="postgres://USERNAME:PASSWORD@HOST:5432/database?connect_timeout=10`

`go run main.go`
