# Too White Powder

## Deploy

```
cp .example.env .env
docker-compose up -d
```

## Devlopment

```
cp .env .dev.env
```

Modify `.dev.env` to fit your environment needs

Export `.dev.env` environment variables

### Frontend

```
cd frontend
npm install
npm run dev
```

> Node.js and npm is required

### Backend

```
go run .
```

> Golang and Git is required
