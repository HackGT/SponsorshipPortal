# SponsorshipPortal [![Build Status](https://travis-ci.org/HackGT/SponsorshipPortal.svg?branch=master)](https://travis-ci.org/HackGT/SponsorshipPortal)

## Getting Started

You must have a working Go and Node.js installation in order to run this project.

If you are new to developing with tools (or new in general), consider installing with your system's package manager when possible.
Please ensure your language and runtime versions match the ones below.

1. [go version 1.9](https://golang.org) [[install](https://golang.org/doc/install)] - For running the webserver for the backend
Make sure you have a `GOHOME` directory set up
2. [dep](https://golang.github.io/dep/) - Dependency manager for go - install with `go get -u github.com/golang/dep/cmd/dep` after go is working
3. [Node.js version 8.9.4](https://nodejs.org/en/) [[download](https://nodejs.org/en/download/)] - Used to build and run the frontend in development
4. [yarn](https://yarnpkg.com/) [[install](https://yarnpkg.com/en/docs/install)] - Dependency manager for Node.js (use `npm` at your own risk!)
5. [Postgres](http://postgresql.org/) - Relational database used as our primary datastore

_Optional_: If you are working on the backend, install
[gin](https://github.com/codegangsta/gin) (`go get github.com/codegangsta/gin`) for live-reloading capabilities.

### On MacOS (with [homebrew](https://brew.sh))

This is a (semi-opinionated) guide to getting started with these tools and languages on MacOS.

> "Package managers are great! Use them!" - Andrew Dai Feb 15 2018

To follow along, setup [homebrew](https://brew.sh) if you have not already.
We will be installing all the tools in the previous section with `brew` or
another package/version manager whenever possible.

```bash
# Install go with homebrew
brew install go
# You might need to setup certain environment variables for go to work.
# Do that now.
# See go's install documentation
# (and also look for resources for go with homebrew and MacOS)

# Install the dep dependency manager for go
go get -u github.com/golang/dep/cmd/dep
# (Optional tool for backend development)
go get github.com/codegangsta/gin

# Check if dep was installed properly (and your go setup)
dep version # Outputs the version and build information

# Install n, a node version manager
brew install n

# Install Node.js 8.9.4 (currently LTS version we are using)
n 8.9.4

# Install yarn with homebrew
brew install yarn --without-node

# Install Postgres.app with homebrew cask (open the App)
brew cask install postgres
# Find and open the Postgres app in your Applications folder
```

[`n`](https://github.com/tj/n) is a version manager for Node.js and makes it very easy to switch versions painlessly.
It is highly recommended to use a tool like this to install and run different versions of Node.js.
For example, `n 8.9.4` installs Node.js version 8.9.4 and sets it as the current version.
You can confirm with `node --version`.

**Note**: `go env` is very helpful in debugging go environment variables.

## Install

Get the source code:
```bash
go get github.com/HackGT/SponsorshipPortal
```
You may get an error `no buildable Go source files`; this is harmless.

For development work run the frontend and backend separately.

### Install dependencies

Before running the frontend for development change the file `frontend/src/js/configs.js` to assign the HOST to `export const HOST = window.location.protocol + '//' + 'localhost:9000';` This needs to be done or server-side communication won't work. (TODO: host should change automatically from environment variable)

Install frontend dependencies:
```bash
# from project root
cd client
yarn
```

Install backend dependencies:
```bash
# from project root
dep ensure
```

## Setup

### Postgres

Make sure that PG_URL contains the connection string SponsorshipPortal will use
to connect to Postgres.

**NOTE:** The golang library `lib/pq`'s `sslmode` defaults to `require` which may cause issues if you do not
have SSL configured. This may not be an issue on `localhost`.

Examples:
```bash
# connect as default user (`postgres`) to default database with sslmode disabled
export PG_URL="postgres://postgres@localhost:5432/?sslmode=disable"

# connect as `portal` user with password `secret`
# to `dev` database with ssl disabled
export PG_URL="postgres://portal:secret@localhost:5432/dev?sslmode=disable"

# same connection as above in an alternate form
# (this will work but not be parsed correctly by `config/config.go`)
export PG_URL="host=localhost port=5432 user=portal password=secret sslmode=disable"
```

More information and examples on [Postgres docs](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING)

### Generate ECDSA Key Pair

The backend is currently configured to use secp521r1. First, download and install OpenSSL if you do not already have it (most flavors of Linux come with OpenSSL):
`sudo apt install openssl`

Navigate to your working directory for the SponsorshipPortal and generate your secp521r1 key pair:
```
openssl -name secp521r1 -out secp521r1.pem
openssl ecparam -in secp521r1.pem -genkey -noout -out ecprivatekey.pem
openssl ec -in ecprivatekey.pem -pubout -out ecpublickey.pem
```

## Start the app!

Run the database migrations (see `migrations/README.md` for more info)
```bash
go run main.go --migrate
# or if you do not want to set PG_URL for the current session
PG_URL="<connection string>" go run main.go --migrate
```

Start the backend
```bash
# from project root
go run main.go
# if you are working on the backend and have gin installed
# (go get github.com/codegangsta/gin)
gin run main.go

# or set PG_URL for the command and run
PG_URL="<connection string>" go run main.go
# also works with `gin`
PG_URL="<connection string>" gin run main.go
```

Then, in a new shell, start the frontend and navigate to `localhost:8500`.
```bash
# from project root
cd frontend
npm start
```

### Building and running

#### Build the front-end

```bash
# build frontend
# from project root...
cd frontend
npm run build
```

#### Build and run the server

```bash
# build backend
# from project root...
go build .

# set connection string
# run migrations and start server
export PG_URL="<connection string>"
./SponsorshipPortal --migrate
./SponsorshipPortal
```

## Contributing

The frontend is a standard React (https://reactjs.org/) app with Redux for state management.

The application is written in go and heavily utilizes the standard library
(`net/http` and `database/sql`) as well as packages from the [gorilla web toolkit](http://www.gorillatoolkit.org)
([`mux`](http://www.gorillatoolkit.org/pkg/mux)).

If you are working only on the backend or the frontend, you can build (and start)
the other part and leave it running while you do your development.
Example: build the frontend and then build and run the backend with `gin`'s live
reloads. See above [instructions](#building-and-running) on how to build each part.

### Code Layout

The directory structure:
```
main.go           Boot script for the server
routes.go         Register non-controller routes go here
server.go         Initialize server and setup components
Gopkg.toml
Gopkg.lock

config/           App server config package

controllers/      App controllers go here
    controller.go Root controller, register other controllers here

database/         Utility package for initializing a database connection from config

logger/           Utility package for initializing a logger from loaded config

models/           App models go here

migrations/       Database migrations (sql scripts) go here

client/           Front-end project root
    package.json
    yarn.lock
    src/          Source code for React app
    static/       Public static assets - this is not used in development
    node_modules/ Dependencies, do not edit this directory - managed by yarn

vendor/           Dependencies, do not edit this directory - managed by dep
```
