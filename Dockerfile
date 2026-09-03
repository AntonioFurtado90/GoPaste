FROM golang:alpine AS dev

WORKDIR /app

CMD ["go", "run", "."]
