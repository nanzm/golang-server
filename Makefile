.PHONY:

# --- mysql redis nsq  ---------------------------------------------------------
base-up:
	cd deployments  && docker-compose -f ./docker-compose.base.yml up -d

base-down:
	cd deployments  && docker-compose -f ./docker-compose.base.yml down

base-logs:
	cd deployments  && docker-compose -f ./docker-compose.base.yml logs

base-ps:
	cd deployments  && docker-compose -f ./docker-compose.base.yml ps

data-clean:
	cd deployments  && rm -rf redis/data && rm -rf mysql/data


# --- app ---------------------------------------------------------
build-transit:
	go build -o transit cmd/transit/main.go

build-manage:
	go build -o manage cmd/manage/main.go

# --- app ---------------------------------------------------------
docker-transit:
	docker build -f ./build/transit.Dockerfile -t nancode/dora-transit .

docker-manage:
	docker build -f ./build/manage.Dockerfile -t nancode/dora-manage .

