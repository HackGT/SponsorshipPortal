FROM alpine

ADD backend.tar.gz /www

CMD /www/run.sh
