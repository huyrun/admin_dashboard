# Admin Dashboard

This repository contains the source code for the Admin Dashboard application, designed for managing administrative tasks with an intuitive interface. The app is containerized using Docker, providing easy setup and deployment.

## Features
- Simple deployment with Docker Compose
- Pre-configured PostgreSQL database
- Scalable and extensible architecture

## Prerequisites
Ensure the following tools are installed:
- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

To run the application locally:

1. Clone the repository:
    ```bash
    git clone https://github.com/huyrun/admin_dashboard.git
    cd admin_dashboard
    ```

2. Start the app with Docker Compose:
    ```bash
    docker-compose -f docker/docker-compose.yml up --build -d
    ```

This will build and start the application in detached mode.

## Change environments

I use a config template to manage environment configurations: [config_template.yml](docker/config_template.yml)

```yaml
<config_name>: {{ default "<default value>" .Env.<ENV_NAME> }}
```

During Docker image build, dockerize parses the template into the config file [config.yml](etc/config/config.yml), which the app uses.

To update environment configurations, modify [config_template.yml](docker/config_template.yml) and run the following command for local testing (no need to run this during Docker image builds):

```shell
dockerize -template docker/config_template.yml:etc/config/config.yml
```

Or, you can directly edit the environment in [config.yml](etc/config/config.yml). Note that this change won't affect Docker image builds unless you also update [config_template.yml](docker/config_template.yml).