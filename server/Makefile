build:
	GOOS=linux go build -o app
	docker build -t gcr.io/sanjaesan/say .
	rm -f app

push: 
	gcloud docker -- push gcr.io/sanjaesan/say