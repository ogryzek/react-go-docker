FROM node:10-jessie AS app_builder
# RUN mkdir /myapp
ADD ./api ./api
ADD ./app ./app

RUN echo "ls /myapp/api $( ls./api )"
RUN echo "ls /myapp/app $( ls ./app )"

WORKDIR ./app
RUN echo "Running from WORKDIR: $( pwd )"
RUN echo "The files here are: $( ls )"

RUN npm install
RUN npm run build

FROM golang:1.12-alpine AS api_builder
COPY --from=app_builder ./api .
COPY --from=app_builder ./app/build/ static/
EXPOSE 3000
CMD go run .
