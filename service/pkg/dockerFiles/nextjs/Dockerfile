FROM node:14

WORKDIR /app

COPY . ./

RUN npm install

EXPOSE 3000

CMD ["sh", "-c", "npm start || npm run dev"]
