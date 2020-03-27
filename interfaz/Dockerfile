FROM node:6.11.1
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
RUN npm install

EXPOSE 8080
CMD [ "npm", "run", "serve" ]
