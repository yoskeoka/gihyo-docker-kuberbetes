FROM alpine:3.7

ENV VERSION=0.11.5

RUN apk --no-cache add --virtual=build-deps curl gnupg

# terraformとチェックサム情報を取得
ADD https://releases.hashicorp.com/terraform/$VERSION/terraform_${VERSION}_linux_amd64.zip .
ADD https://releases.hashicorp.com/terraform/$VERSION/terraform_${VERSION}_SHA256SUMS .
ADD https://releases.hashicorp.com/terraform/$VERSION/terraform_${VERSION}_SHA256SUMS.sig .

# Keybaseで公開されているHashicorpのPGP公開鍵をインポートする
RUN curl https://keybase.io/hashicorp/pgp_keys.asc | gpg --import

# チェックサムの署名を検証する
RUN gpg --verify terraform_${VERSION}_SHA256SUMS.sig terraform_${VERSION}_SHA256SUMS

# アーカイブファイルのチェックサムを検証する
RUN cat terraform_${VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -cs

RUN unzip terraform_${VERSION}_linux_amd64.zip
RUN mv terraform /usr/local/bin/
RUN terraform -v
