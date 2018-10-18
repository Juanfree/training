FROM docker.elastic.co/elasticsearch/elasticsearch:6.4.1

ADD docker/rootfs /

RUN ./bin/elasticsearch-plugin install https://github.com/NLPchina/elasticsearch-sql/releases/download/6.4.1.0/elasticsearch-sql-6.4.1.0.zip
