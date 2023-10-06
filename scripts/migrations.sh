#!/bin/sh
domain_kw=education
subdomain=person

cd migrations
set -o allexport; eval $(cat .env); set +o allexport

case $1 in
up)
  goose postgres "host=localhost sslmode=disable dbname=${domain_kw}_${subdomain}_api port=5432" up $2
  ;;
down)
  goose postgres "host=localhost sslmode=disable dbname=${domain_kw}_${subdomain}_api port=5432" down $2
  ;;
status)
  ;;
*)
  echo "Usage: $(basename $0) {up|down} [serial]"
  ;;
esac

goose postgres "host=localhost sslmode=disable dbname=${domain_kw}_${subdomain}_api port=5432" status
