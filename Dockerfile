# rsyslog redhat ubi9 with daemon 
# YOU HAVE TO RUN IT IN DOCKER CONTAINER PRIVILIDGE MODE
FROM registry.access.redhat.com/ubi9-init

RUN dnf install rsyslog -y
RUN dnf install net-tools -y
RUN printf "\nmodule(load=\"imtcp\")\ninput(type=\"imtcp\" port=\"10123\")\n" >> /etc/rsyslog.conf
RUN printf "\n*.* /var/log/mytest.all\n" >> /etc/rsyslog.conf
RUN printf "\n*.info /var/log/mytest.info\n" >> /etc/rsyslog.conf
RUN printf "\n*.err /var/log/mytest.err\n" >> /etc/rsyslog.conf
RUN printf "\n*.crit /var/log/mytest.crit\n" >> /etc/rsyslog.conf