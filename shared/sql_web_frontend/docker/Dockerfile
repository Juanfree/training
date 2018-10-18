FROM node:alpine

ADD docker/rootfs /

WORKDIR /opt/app/site-server

RUN npm install --save

ENTRYPOINT node node-server.js
