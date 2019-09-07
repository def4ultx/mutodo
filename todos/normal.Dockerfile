FROM golang:1.13

# if left blank app will run with dev settings
ARG app_env
ENV APP_ENV $app_env

COPY . /todos

WORKDIR /todos

RUN go mod download
RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	fresh -c fresh_ignore.conf; \
fi

EXPOSE 8080