STRIPE_SECRET=
STRIPE_KEY=
GOSTRIPE_PORT=4000
API_PORT=4001

clean:
	rm -r dist
	go clean

build_web: clean
	@go build -o dist/web ./cmd/web

web: build_web
	@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} ./dist/web -port=${GOSTRIPE_PORT} &

build_api: clean
	@go build -o dist/api ./cmd/api

api: build_api
	@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET}  ./dist/api -port=${API_PORT} &

build: build_web build_api

app: web api

stop_web:
	@echo "Stopping the front end..."
	@-pkill -SIGTERM -f "web -port=${GOSTRIPE_PORT}"
	@echo "Stopped front end"

stop_api:
	@echo "Stopping the back end..."
	@-pkill -SIGTERM -f "api -port=${API_PORT}"
	@echo "Stopped back end"

stop_app: stop_web stop_api


.PHONY=web api clean build build_api build_web app stop_api stop_web stop_app