FROM centos
MAINTAINER mingz2013
COPY chinese-chess-table-go /sbin/chinese-chess-table-go
RUN chmod 755 /sbin/chinese-chess-table-go
ENTRYPOINT ["/sbin/chinese-chess-table-go"]
