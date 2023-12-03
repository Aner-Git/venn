# Yarn comes bundled with the official node docker image. Nice.
FROM node:20-bullseye-slim as local

ADD docker/react-app/rewrite-index.sh /rewrite-index.sh
RUN chmod +x /rewrite-index.sh

# Configure working dir
RUN mkdir /webapp
VOLUME ["/webapp"]
WORKDIR /webapp

CMD ["yarn", "start:local-docker"]

