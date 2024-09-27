FROM postgres:latest


USER root
RUN apt-get update && apt-get install -y \
  build-essential \
  git \
  postgresql-server-dev-all \
  && rm -rf /var/lib/apt/lists/*


RUN cd /tmp \
  && git clone --branch v0.7.4 https://github.com/pgvector/pgvector.git \
  && cd pgvector \
  && make \
  && make install