SHELL=/bin/zsh

firebase-emulator/start:
	firebase emulators:start --import=./data
firebase-emulator/export:
	firebase emulators:export ./data
watch:
	cd api;air
