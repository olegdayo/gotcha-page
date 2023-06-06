FROM node:latest
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY package-lock.json ./
RUN npm install react-scripts@3.4.1 -g

COPY . .

RUN npm run build
RUN npm install -g serve --save
CMD ["serve", "-s", "build"]
