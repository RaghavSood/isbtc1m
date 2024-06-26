# Use the NixOS base image
FROM nixos/nix as builder

# Set up the Nix environment
COPY . /src
WORKDIR /src

RUN nix \
    --extra-experimental-features "nix-command flakes" \
    --option filter-syscalls false \
    build

RUN mkdir -p /tmp/nix-store-closure /tmp/litestream
RUN cp -R $(nix-store -qR result/) /tmp/nix-store-closure

FROM alpine:3 as alpine

RUN apk add -U --no-cache ca-certificates

WORKDIR /app

# Copy /nix/store
COPY --from=builder /tmp/nix-store-closure /nix/store
COPY --from=builder /src/result /app
CMD [ "/app/bin/isbtc1m" ]
