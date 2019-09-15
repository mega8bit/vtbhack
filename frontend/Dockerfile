FROM node:alpine as builder
WORKDIR /app
COPY package.json package.json
COPY yarn.lock yarn.lock
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.14.1-alpine
COPY default.conf /etc/nginx/conf.d/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
