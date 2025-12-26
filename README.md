# migration service:
- It will init database tables and required entities.
- It will dump data from migration/newsdump.json file or can do it from API for latest news.
- It will connect with LLM to add summary for news data dump.
- For LLM connection, Setup gcloud application default credentials in local at path ~/.config/gcloud/application_default_credentials.json to use google api's.

# news api service:
- `category API` - Get latest news for categoryName provided.

  `curl --location 'http://localhost:8080/api/v1/news/category?categoryName=world&page=1&limit=10'`
  <img width="1049" height="745" alt="Screenshot 2025-12-26 at 2 57 44 PM" src="https://github.com/user-attachments/assets/032e5458-b0e8-4f48-a841-5bb9b01098a0" />

- `source API` - Get latest news for source provided.

  `curl --location 'http://localhost:8080/api/v1/news/source?sources=world&page=1&limit=10'`
  <img width="1129" height="740" alt="Screenshot 2025-12-26 at 2 59 06 PM" src="https://github.com/user-attachments/assets/94e88eb8-9b7a-4983-9be7-665086d158c8" />

- `score API` - Get news rank by relevance score (highest first).

  `curl --location 'http://localhost:8080/api/v1/news/score?relevanceScore=world&page=1&limit=10'`
  <img width="1097" height="734" alt="Screenshot 2025-12-26 at 2 59 54 PM" src="https://github.com/user-attachments/assets/90b9683f-9f7e-4baa-92ef-1e2500590cbf" />

- `search API` - Get news rank by relevance score (highest first).

  `curl --location 'http://localhost:8080/api/v1/news/search?query=elon%20musk&page=3&limit=2'`
  <img width="1100" height="724" alt="Screenshot 2025-12-26 at 3 01 09 PM" src="https://github.com/user-attachments/assets/6ad6c4ae-5816-42aa-9fef-a5e19db4316e" />


# To start services.
1. Create a project in GCloud, to user google vertex api for summary generation for news.
2. Enable Vertex AI API in GCloud
3. Use project ID in env for key GOOGLE_CLOUD_PROJECT.
4. Migration service will run before API service.
5. Update env file as required.
6. Run docker compose build 
7. Run docker compose up
   
# 🐳 Docker Setup
docker-compose.yml (Overview)
1. postgres
2. migration-service
3. news-api-service

   `Migration runs before API.`

# 🌱 Environment Variables
.env example
```
# App
APP_ENV=local
APP_PORT=8080

CREATE_TABLES=true
API_DATA_MIGRATION=true
FILE_DATA_MIGRATION=true

# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=psql
DB_NAME=news_db

GOOGLE_CLOUD_PROJECT={projectid}
GOOGLE_CLOUD_LOCATION=global
GOOGLE_GENAI_USE_VERTEXAI=True
```

# 🔐 Google Vertex AI Setup
1. Create a GCP project
2. Enable Vertex AI API
3. Generate Application Default Credentials
4. Place credentials at:
    ```
   ~/.config/gcloud/application_default_credentials.json
    ```
6. Mount credentials into container

# 🚀 How to Run
1. docker compose build
2. docker compose up


`Migration runs first, API starts after.`
