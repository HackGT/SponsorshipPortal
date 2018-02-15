# SponsorshipPortal
[![Build Status](https://travis-ci.org/HackGT/SponsorshipPortal.svg?branch=master)](https://travis-ci.org/HackGT/SponsorshipPortal)

## Getting Started

You must have a working Go and Node.js installation in order to run this project.
- Instructions to install Go can be found here: https://golang.org/doc/install. Make sure you have a `GOHOME` directory set up.
- Install Node.js from the official site: https://nodejs.org/en/download.

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
npm install
```

Install backend dependencies:
```
# from project root
cd backend
go get github.com/revel/cmd/revel
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

### Start the app!

Start the backend
```
# from project root
cd backend
revel run
```

Start the frontend and navigate to `localhost:8500`.
```
# from project root
cd frontend
npm start
```

## Contributing

The frontend is a standard React (https://reactjs.org/) app with Redux for state management.

The backend is written using the Revel framework. Please read the tutorial on the official site: http://revel.github.io/.

## Backend Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites
