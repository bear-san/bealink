FROM node:20.12 AS builder

WORKDIR /app
ADD package.json package-lock.json /app/
RUN npm install

ADD . /app
RUN npm run build

FROM nginx:1.26
COPY --from=builder /app/out /usr/share/nginx/html

EXPOSE 80
