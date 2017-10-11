FROM debian:8

ADD backend.tar.gz /www

RUN apt-get update && apt-get install -y \
    ruby \
    ruby-dev \
    build-essential \
    default-jre

RUN gem install yomu

CMD /www/run.sh
