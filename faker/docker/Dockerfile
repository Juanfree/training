FROM python:latest

WORKDIR /opt/faker

ENV ELASTIC_SEARCH_URL=http://elastic_search:9200
ENV NUMBER_OF_DOCS=1000
ENV INDEX_NAME=training
ENV FORCE_INIT_INDEX=True
ENV FORMAT=company_id:int,company_name:str,location:str,transactions:int,last_updated:ts

ADD docker/rootfs /

RUN pip install virtualenv

RUN virtualenv -p /usr/bin/python2.7 venv

RUN venv/bin/pip install -r requirements.txt

ENTRYPOINT ./venv/bin/python es_test_data.py \
    --es_url=$ELASTIC_SEARCH_URL \
    --count=$NUMBER_OF_DOCS \
    --force_init_index=$FORCE_INIT_INDEX \
    --index_name=$INDEX_NAME \
    --format=$FORMAT
