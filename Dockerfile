FROM debian:8

ADD backend.tar.gz /www

CMD /www/run.sh
