FROM node:latest

WORKDIR /app

COPY package.json package-lock.json /app/

RUN npm install

COPY . .

RUN npx vite build

RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "dist"]
