# go-workspace
golang programs from go.dev

Gorilla's mux package, requests can be matched based on URL host, path, path prefix, schemes, header and query values, and HTTP methods.

go get github.com/gorilla/mux

Mongo DB :
----------
docker pull mongo
docker image inspect mongo
docker run -d  --name mongo-on-docker  -p 27888:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo

#created folder to use for storing MongoDB data
sudo mkdir /var/dbdata
docker run -it -d -v /var/dbdata:/data/db -p 27017:27017 --name mongodb mongo
docker start mongodb
sudo docker exec -it mongodb bash

docker run --link mongo_db_name_container:mongo -p 8081:8081 -e ME_CONFIG_MONGODB_URL="mongodb://mongo:27017" mongo-express
docker run --link mongodb:mongo -p 8081:8081 -e ME_CONFIG_MONGODB_URL="mongodb://mongo:27017" mongo-express

mongo ui -- http://localhost:8081


mongodb://<username>:<password>@<host>:<port>/?authSource=admin
mongodb://mongoadmin:secret@localhost:27888/?authSource=admin


https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial