# Use an official Golang runtime as a parent image
FROM node:8.11

# Set the working directory to golang working space
WORKDIR /app
RUN npm install http-server -g
# Copy the current directory contents into the container at current directory
ADD localhost:6060 server

CMD http-server server -p 8080