# graphql-tutorial

## Get Started

**Environment variables.**

```
% cp .env.example .env
```

Now add your Auth0 and Database variables.

**Install dependencies and run the server**

```
% go mod download
% make run
```

## Commands

```
% make run # starts the server
% make token # generates an access token with auth0
% make db # starts a docker container with mysql database
% make db-console # gets access to the docker container
% make generate # generates new graphsql models and resolvers
% make migration name=<NAME> # creates a new migration
```
