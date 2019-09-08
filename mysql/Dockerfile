FROM mysql:8

COPY sql/ /docker-entrypoint-initdb.d/

RUN mkdir logs
CMD ["mysqld"]