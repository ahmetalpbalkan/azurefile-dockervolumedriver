FROM alpine:3.6

RUN apk --update add ca-certificates cifs-utils
RUN mkdir -p /run/docker/plugins /var/run/docker/volumedriver/azurefile

ADD azurefile-dockervolumedriver /usr/bin/azurefile-dockervolumedriver
# ADD start.sh /start.sh

# ENTRYPOINT [ "/start.sh" ]
CMD [ "azurefile-dockervolumedriver" ]