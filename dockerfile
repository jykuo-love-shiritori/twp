FROM node:20.7.0
WORKDIR /home/user
COPY ./frontend .
RUN npm install
RUN npm run build 
RUN npm run export
CMD [ "npm"]





# Use the official Golang image as the base image
FROM golang:1.21.1
# Set the working directory inside the container
WORKDIR /home/user/


# COPY --from=0 /home/user/out /home/user/build
COPY --from=0 /home/user/ /home/user/frontend
# Copy the GoLang application source code into the container
COPY ./*.go ./pkg ./go.mod ./

# RUN go mod init main

# RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware
# Build the GoLang application
RUN go build -o main
# Command to run your GoLang application
CMD ["./main"]
