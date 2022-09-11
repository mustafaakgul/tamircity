# Tamircity

# TODO
- Include swagger.json,.yaml in docs folder

# Database ER Diagram

# API Response Model

# API Response Model
1. Create DB Manually SHELL
2. Migration SHELL
3. Seed SHELL

# Docker Local
1. docker-compose -f docker-compose-common.yml -f docker-compose-local-dev.yml up --build

# Docker Production
1. docker-compose -f docker-compose-common.yml -f docker-compose-prod.yml up --build

# Docker Development
1. docker build -t tamircity-development .
2. docker images
3. docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/tamircity tamircity-development

# Docker Production
1. docker build -t mathapp-production -f Dockerfile.production .
2. docker run -it -p 8010:8010 mathapp-production

# Enabling CI / CD
1. Automatically build
2. Automatically test
3. Automatically linting
2. Autimatically deploy