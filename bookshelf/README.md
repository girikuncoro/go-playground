This bookshelf package is originally created in Google to provide samples running app on Google Cloud Platform. I port this by myself to brush up my golang and commandline skill.

## How to run
Run a mysql instance, you can use docker
```sh
docker run --name <DB_HOST> -p <DB_PORT>:3306 -e MYSQL_ROOT_PASSWORD=<DB_PASSWORD> -d mysql:8
```

example:
```sh
docker run --name giri-mysql -p 9000:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8
```

Validate your new instance by accessing it from mysql client:
```sh
docker run -it --link <DB_HOST>:mysql --rm mysql sh -c 'exec mysql -h<DB_HOST> -P3306 -uroot -p<DB_PASSWORD>'
```

example:
```sh
docker run -it --link giri-mysql:mysql --rm mysql sh -c 'exec mysql -hgiri-mysql -P3306 -uroot -pmy-secret-pw'
```

## Test your CLI
You can test CLI without building it by running `go run`

Add book example:
```sh
go run main.go add book -t "Zen Art" -a "Steve Jobs" -d "10 Jan 2004" --host 0.0.0.0 --user root --password my-secret-pw --port 9000
```

List books example:
```sh
go run main.go list book --host 0.0.0.0 --user root --password my-secret-pw --port 9000
```
