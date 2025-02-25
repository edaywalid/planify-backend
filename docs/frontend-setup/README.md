# Backend Setup for Frontend Developers

## Pre-Built Docker Image

To run the backend service, you don’t need to build or manage it manually. Follow the steps below:

1. Install Docker on your machine (if not already installed).
2. Create a docker-compose.yml file
```yml
services:
  backend:
    container_name: planify-backend
    image: imewalid/planify-back:latest
    restart: unless-stopped
    depends_on:
      psql_database:
        condition: service_healthy
    ports:
      - "${PORT}:8080"
    networks:
      - app-network
    env_file:
      - .env

  psql_database:
    container_name: psql_database
    image: postgres:16.3
    restart: always
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: password
    ports:
      - "${HOST_PSQL_PORT}:5432"
    volumes:
      - psqldb:/var/lib/postgresql/data
    networks:
      - app-network
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${DB_USER} -d ${DB_NAME}"]
      interval: 5s
      timeout: 5s
      retries: 5

volumes:
  psqldb:

networks:
  planify-back-network:
    driver: bridge
```
3. Create `.env` file following `./.env.example`

4. Run the backend container:
   ```bash
   docker-compose up -d 
   ```
5. Access the backend at `http://localhost:${PORT}` (or update the API URL in the frontend accordingly).

That’s it! Let us know if you encounter any issues.

