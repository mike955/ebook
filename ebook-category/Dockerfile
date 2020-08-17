FROM node:14.8.0-alpine3.12

WORKDIR /root/app
COPY ./ /root/app

RUN npm config set registry 'https://registry.npm.taobao.org' \
  && npm install \
  && npx tsc 

ENTRYPOINT [ "npm start" ]
