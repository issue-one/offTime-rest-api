# offTime REST API 

Built using [Go Swagger](https://goswagger.io/) and [gORM](https://gorm.io)

Use the command in `generate-server.sh` to regenerate the server.

Use `run.sh` (or `run.bat` if on Windows) to run the server.

Accesss `/docs` route for SwaggerUI.

Accepts the following enviroment variables:

```ini
PORT=8080
IMAGE_STORAGE_PATH=data/images/
IMAGE_SERVING_ROUTE=/images/
```