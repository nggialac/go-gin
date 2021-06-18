# Start from the lastest golang base image
FROM golang:latest

#Add Maintainer Info
LABEL maintainer="Quique <nggialac99@gmail.com>"

#Set the Current working directory inside the container
WORKDIR /app

#Copy Go modules dependency requirements file
COPY go.mod .

#Copy Go Modules expected hashes file
COPY go.sum .

#Download dependencies
RUN go mod download

#Copy all the app sources (recursively copies files and directories from the host into )
COPY . .

#Set http port
ENV PORT 5000

#Build the app
RUN go build

#Remove source files
RUN find . -name "*.go" -type f -delete

#Make port 5000 available to the world outside this container
EXPOSE $PORT

#Run the app
CMD ["./go-gin"]