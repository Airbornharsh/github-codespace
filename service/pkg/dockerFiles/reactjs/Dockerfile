# pull official base image
FROM node:20.5.0-alpine3.18

WORKDIR /app

COPY . ./
RUN npm install

EXPOSE 3000

CMD ["sh", "-c", "npm run dev || npm start"]

