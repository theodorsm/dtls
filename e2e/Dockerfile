# SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
# SPDX-License-Identifier: MIT

FROM docker.io/library/golang:1.18-bullseye

COPY . /go/src/github.com/theodorsm/dtls
WORKDIR /go/src/github.com/theodorsm/dtls/e2e

CMD ["go", "test", "-tags=openssl", "-v", "."]
