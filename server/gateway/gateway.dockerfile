### BASE GO IMAGE
FROM golang:1.22-alpine as builder

# creates app folder
RUN mkdir /app

# copies everything from the root folder into app folder
COPY ./ /app

# sets app folder as working directory for next instructions
WORKDIR /app

# builds the gateway service executable into the previously set working directory (app folder)
RUN CGO_ENABLED=0 go build -o gateway .

# adds the executable rights to gateway executable
RUN chmod +x /app/gateway



### BUILD A TINY DOCKER IMAGE
FROM alpine:latest

# creates app folder
RUN mkdir /app

# copies the previously created builder image from /app/gateway of the previous image into /app of the current image 
COPY --from=builder /app/gateway /app

# runs the executable
CMD ["/app/gateway"]