FROM iron/go
ENV GIN_MODE=release
WORKDIR /app
ADD . /app/
EXPOSE 8080
ENTRYPOINT ["./app"]