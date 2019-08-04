FROM node:10-jessie AS app_builder
ADD app/ .
RUN npm install
RUN npm run build

FROM golang:1.12-alpine
ADD api/ .
COPY --from=app_builder build/ static/
EXPOSE 3000
CMD go run .
