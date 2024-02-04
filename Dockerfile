FROM --platform=linux/amd64 archlinux:base-devel

WORKDIR /mnt

RUN pacman-key --init
RUN pacman -Sy --noconfirm reflector arch-install-scripts

RUN reflector --latest 5 \
  --protocol http,https \
  --save "/etc/pacman.d/mirrorlist" \
  --sort rate

RUN pacman -Syu --noconfirm base base-devel

COPY --chown=user:group ./dist /mnt

CMD [ "./app" ]
