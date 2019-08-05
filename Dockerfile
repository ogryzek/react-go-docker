FROM node:10-jessie AS app_builder
ADD . .

WORKDIR app
RUN npm install 
RUN npm run build 

FROM golang:1.12-alpine AS api_builder
COPY --from=app_builder ./api .
COPY --from=app_builder app/build/ static/
ENV PORT=3000
EXPOSE ${PORT}
CMD go run .
