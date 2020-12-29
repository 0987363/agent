FROM heifeng/alpine-node-python3
MAINTAINER heifeng


WORKDIR /app

COPY /agent /app
COPY py py/

EXPOSE 80
CMD /app/agent 
