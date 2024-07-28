compile-proto:
	trpc create \
		-p app.proto \
		-d ./src \
		-f \
		-o ./client/_generated \
		--nogomod \
		--mock=false \
