FROM node:20-slim
WORKDIR /app

COPY package*.json ./

RUN npm install

# 3. Copy the rest
COPY . .

EXPOSE 5173
EXPOSE 6006
