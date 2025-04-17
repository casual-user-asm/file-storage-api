
# File storage service

A backend service for uploading and retrieving files using MinIO (S3-compatible storage).


## Installation

1 . Clone the repository:

```bash
  git clone https://github.com/casual-user-asm/file-storage-service.git
  cd file-storage-service
```

2 . Build and run the Docker containers:

```bash
  docker-compose up -d
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`MINIO_ROOT_USER`

`MINIO_ROOT_PASSWORD`

`MINIO_ENDPOINT`

`MINIO_BUCKET`
