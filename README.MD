# SponsorshipPortal [![Build Status](https://travis-ci.org/HackGT/SponsorshipPortal.svg?branch=master)](https://travis-ci.org/HackGT/SponsorshipPortal)

## Getting Started

You must have a working Go and Node.js installation in order to run this project. 

Instructions to install Go can be found here: https://golang.org/doc/install. Make sure you have a GOHOME directory setup.

Install Node.js from the official site: https://nodejs.org/en/download/.

## Installing

Get the source code by executing `go get github.com/HackGT/SponsorshipPortal` 

You may get an error `no buildable Go source files`, this is harmless.

For development work run the frontend and backend separately.

Before running the frontend for development change the file `frontend/src/js/configs.js` to assign the HOST to `export const HOST = window.location.protocol + '//' + 'localhost:9000';` This needs to be done or server-side communication won't work. (TODO: host should change automatically from environment variable)

To run the frontend navigate to the `frontend` directory and execute 

```
npm install
npm start
```

Navigate to the `backend` directory and execute `go get github.com/revel/cmd/revel`

You should now be ready to run the backend. Execute `revel run`

## Contributing

The frontend is a standard React app with Redux for state management. React documentation can be found here: https://reactjs.org/.

The backend is written using the Revel framework. Please read the tutorial on the official site: http://revel.github.io/.
