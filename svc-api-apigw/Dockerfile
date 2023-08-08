FROM centos:7

ADD ../istio /

RUN chmod 777 /istio \
&& rm -f /etc/localtime \
&& ln -sv /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/istio $PARAMS"]
