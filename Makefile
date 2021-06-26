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


# --- elasticsearch ---------------------------------------------------------
elastic-up:
	cd deployments  && docker-compose -f ./elasticstack.yml up -d

elastic-down:
	cd deployments  && docker-compose -f ./elasticstack.yml down

elastic-logs:
	cd deployments  && docker-compose -f ./elasticstack.yml logs


# --- app ---------------------------------------------------------
dev-transit:
	go run cmd/transit/main.go

dev-manage:
	go run cmd/manage/main.go


# --- app ---------------------------------------------------------
build-transit:
	docker build -f deployments/transit.Dockerfile -t nancode/dora-transit

build-manage:
	docker build -f deployments/manage.Dockerfile -t nancode/dora-manage

