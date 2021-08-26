electron:
	cd electron_app && yarn start

ui:
	yarn --cwd ./frontend dev

run:
	go run ./backend -mode=dev


install: 
	yarn --cwd ./frontend install