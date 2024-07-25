build:
	docker compose build
run:
	docker compose up -d
push:
	docker login --username ${GITHUB_USERNAME} --password ${GITHUB_TOKEN} ghcr.io
	docker push ghcr.io/${GITHUB_USERNAME}/consumer:1.0
	docker push ghcr.io/${GITHUB_USERNAME}/api_kafka_sarama:1.0