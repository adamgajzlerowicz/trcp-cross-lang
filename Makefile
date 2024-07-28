compile-proto:
	trpc create \
		-p app.proto \
		-d ./src \
		-o ./client/_generated \
		--nogomod \
		--mock=false \
