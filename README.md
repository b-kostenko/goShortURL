Connect to database:
```bash
docker run -d \
  --name shorturl-postgres \
  --env-file .env \
  -p 5432:5432 \
  postgres:latest
```

## 📄 Environment Variables
Create a `.env` file using the sample provided:

```bash
cp .env.sample .env
```