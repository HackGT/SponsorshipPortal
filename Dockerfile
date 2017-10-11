FROM debian:8

ADD backend.tar.gz /www

RUN apt-get update && apt-get install ruby-full && apt-get install default-jre && gem install yomu

CMD /www/run.sh
