FROM node:10-jessie AS app_builder
RUN echo "Starting in: $(pwd)"
RUN echo "I am:  $( whoami )"
RUN echo "files are : $( ls -a )"

ADD app/ /app
WORKDIR /app

RUN echo "Changed WORKDIR: $WORKDIR is it: $(pwd)"
RUN echo "I am now:  $( whoami )"
RUN echo "files here are are : $( ls -a )"

RUN npm install
RUN npm run build

FROM golang:1.12-alpine
ADD api/ .
COPY --from=app_builder /app/build/ static/
EXPOSE 3000
CMD go run .