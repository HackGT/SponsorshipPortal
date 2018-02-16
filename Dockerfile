FROM node:carbon-alpine as build-frontend
WORKDIR /frontend
COPY frontend/package.json ./
COPY frontend/yarn.lock ./
RUN yarn --production
COPY frontend .
RUN npm run build

FROM golang:1.9-alpine as build-backend
WORKDIR /go/src/github.com/HackGT/SponsorshipPortal/backend
RUN apk update && apk add git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/revel/cmd/revel
COPY backend/Gopkg.* ./
RUN dep ensure -vendor-only
COPY backend .
COPY --from=build-frontend frontend/build ./app/public
RUN revel package github.com/HackGT/SponsorshipPortal/backend prod
RUN pwd
RUN ls

FROM alpine:latest as run-server
EXPOSE 9000
WORKDIR /www
COPY --from=build-backend /go/src/github.com/HackGT/SponsorshipPortal/backend/backend.tar.gz .
RUN tar -xf backend.tar.gz
CMD sh run.sh