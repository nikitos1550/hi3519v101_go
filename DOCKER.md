# DOCKER notes

```
FROM debian

RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get -y install \
    build-essential cmake gawk libncurses-dev libc6-dev intltool python \
    git subversion time unzip zlib1g-dev libssl-dev wget curl cpio bc \
    gettext gettext-base liblocale-gettext-perl upx \
    flex texinfo help2man libtool-bin byacc bison pkg-config libyaml-dev

RUN git clone --recursive https://github.com/OpenHisiIpCam/br-hisicam --depth 1

WORKDIR ./br-hisicam
RUN make install-ubuntu-deps
RUN make prepare
RUN make list-configs
```
