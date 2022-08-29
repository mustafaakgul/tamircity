#!/usr/bin/env bash
docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=tamir -p 5432:5432 -d postgres
