electron:
	cd electron_app && yarn start

ui:
	cd frontend && yarn dev

run:
	go run ./backend -mode=dev