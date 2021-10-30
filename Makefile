firebase-emulator/start:
	firebase emulators:start --import=./data
firebase-emulator/export:
	firebase emulators:export ./data
build/for-server-dev:
	go run api/app.go
