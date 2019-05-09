
build:
	docker build --build-arg GIT_COMMIT=$(shell git rev-list -1 HEAD) -t kfbi/budget-discord-nitro-api:latest budgetDiscordNitro/
	docker push kfbi/budget-discord-nitro-api:latest



