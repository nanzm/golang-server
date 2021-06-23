.PHONY:

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

build-transit:
	docker build -f deployments/transit.Dockerfile -t nancode/dora-transit

build-manage:
	docker build -f deployments/manage.Dockerfile -t nancode/dora-manage

