FROM fedora:34 

RUN dnf update -y && dnf install -y virt-install && dnf clean all
COPY ./updater /usr/bin/updater

ENTRYPOINT ["/usr/bin/updater"]
