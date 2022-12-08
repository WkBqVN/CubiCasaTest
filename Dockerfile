# ## We specify the base image we need for our
# ## go application
FROM golang:1.19
# ## We create an /app directory within our
# ## image that will hold our application source
# ## files
RUN mkdir /app
# ## We copy everything in the root directory
# ## into our /app directory
ADD . /app
# ## We specify that we now wish to execute
# ## any further commands inside our /app
# ## directory
WORKDIR /app

RUN go mod download

# ## we run go build to compile the binary
# ## executable of our Go program
COPY . .
COPY .env .

RUN go build -o /main

EXPOSE 8080
# Command to run when starting the container
CMD ["/main"]
# ## Our start command which kicks off
# ## our newly created binary executable
