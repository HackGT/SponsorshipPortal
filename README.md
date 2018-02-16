# SponsorshipPortal [![Build Status](https://travis-ci.org/HackGT/SponsorshipPortal.svg?branch=master)](https://travis-ci.org/HackGT/SponsorshipPortal)

## Getting Started

You must have a working Go and Node.js installation in order to run this project.
- Instructions to install Go can be found here: https://golang.org/doc/install.
  - After Go is running, install `dep`, the dependency manager we are using for this project. `go get -u github.com/golang/dep/cmd/dep`
- Install Node.js from the official site: https://nodejs.org/en/download.
  - This project uses the `yarn` dependency manager. Find more information at https://yarnpkg.com. You can use `npm` at your own risk.

If you are new to developing with tools (or new in general), consider installing with your system's package manager when possible.
Please ensure your language and runtime versions match the ones below.

1. [go 1.9](https://golang.org) - For running the webserver for the backend
  - [install](https://golang.org/doc/install)
  - Make sure you have a `GOHOME` directory set up
2. [dep](https://golang.github.io/dep/) - Dependency manager for go
  - install with `go get -u github.com/golang/dep/cmd/dep` after go is working
3. [Node.js 8.9.4](https://nodejs.org/en/) - Used to build and run the frontend in development
  - [[download](https://nodejs.org/en/download/)]
4. [yarn](https://yarnpkg.com/) - Dependency manager for Node.js (use `npm` at your own risk!)
  - [[install](https://yarnpkg.com/en/docs/install)]

### On MacOS (with [homebrew](https://brew.sh))

This is a (semi-opinionated) guide to getting started with these tools and languages on MacOS.

> Package managers are great! Use them!
> - Andrew Dai Feb 15 2018

To follow along, setup [homebrew](https://brew.sh) if you have not already.
We will be installing all the tools in the previous section with `brew` or
another package/version manager whenever possible.

```
# Install go with homebrew
brew install go
# You might need to setup certain environment variables for go to work.
# Do that now.
# See go's install documentation
# (and also look for resources for go with homebrew and MacOS)

# Install the dep dependency manager for go
go get -u github.com/golang/dep/cmd/dep

# Check if dep was installed properly (and your go setup)
dep version # Outputs the version and build information

# Install n, a node version manager
brew install n

# Install Node.js 8.9.4 (currently LTS version we are using)
n 8.9.4

# Install yarn with homebrew
brew install yarn --without-node
```

[`n`](https://github.com/tj/n) is a version manager for Node.js and makes it very easy to switch versions painlessly.
It is highly recommended to use a tool like this to install and run different versions of Node.js.
For example, `n 8.9.4` installs Node.js version 8.9.4 and sets it as the current version.
You can confirm with `node --version`.

**Note**: `go env` is very helpful in debugging go environment variables.

## Installing

Get the source code:
```
go get github.com/HackGT/SponsorshipPortal
```
You may get an error `no buildable Go source files`; this is harmless.

For development work run the frontend and backend separately.

### Install dependencies

Before running the frontend for development change the file `frontend/src/js/configs.js` to assign the HOST to `export const HOST = window.location.protocol + '//' + 'localhost:9000';` This needs to be done or server-side communication won't work. (TODO: host should change automatically from environment variable)

Install frontend dependencies:
```
# from project root
cd frontend
yarn
```

Install backend dependencies:
```
# from project root
cd backend
go get -u github.com/revel/cmd/revel
dep ensure
```

### Start the app!

Start the backend
```
# from project root
cd backend
revel run
```

Then, in a new shell, start the frontend and navigate to `localhost:8500`.
```
# from project root
cd frontend
npm start
```

## Contributing

The frontend is a standard React (https://reactjs.org/) app with Redux for state management.

The backend is written using the Revel framework. Please read the tutorial on the official site: http://revel.github.io/.

## Backend Code Layout

The directory structure of a Revel application:
```
conf/             Configuration directory
    app.conf      Main app configuration file
    routes        Routes definition file

app/              App sources
    init.go       Interceptor registration
    controllers/  App controllers go here

public/           Public static assets - this is not used in development

tests/            Test suites

vendor/           Vendored dependencies, do not edit this directory - managed by dep
```
