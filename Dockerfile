FROM node:carbon-alpine as build-frontend
WORKDIR /client
COPY ./client/package.json .
COPY ./client/yarn.lock .
RUN yarn --production
COPY client .
RUN npm run build

FROM golang:1.9-alpine as build-backend
WORKDIR /go/src/github.com/HackGT/SponsorshipPortal
RUN apk update && apk add git
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.* ./
RUN dep ensure -vendor-only
COPY . .
RUN go build

FROM alpine:latest
WORKDIR /www
COPY --from=build-frontend /client/static client/static
COPY --from=build-backend /go/src/github.com/HackGT/SponsorshipPortal .
CMD ./SponsorshipPortal