#!/bin/sh
domain_kw=education
subdomain=person

cd migrations
set -o allexport; eval $(cat .env); set +o allexport
goose postgres "host=localhost sslmode=disable dbname=${domain_kw}_${subdomain}_api port=5432" up
goose postgres "host=localhost sslmode=disable dbname=${domain_kw}_${subdomain}_api port=5432" status
