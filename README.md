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
