FROM alpine:latest
MAINTAINER jerluc <jerluc@drft.rs> @jerluc

# only need ca-certificates & openssl if want to use https_google
RUN apk --update add bind-tools ca-certificates openssl && update-ca-certificates && rm -rf /var/cache/apk/*

ADD dns /dns
ADD Corefile /Corefile

EXPOSE 53 53/udp
ENTRYPOINT ["/dns", "-conf", "/Corefile"]
