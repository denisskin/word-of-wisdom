.PHONY:
docker-build:
	DOCKER_BUILDKIT=0 docker build --no-cache -f build/server/Dockerfile -t wow-server .
	DOCKER_BUILDKIT=0 docker build --no-cache -f build/client/Dockerfile -t wow-client .
