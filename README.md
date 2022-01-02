# go-workspace
golang programs from go.dev

Gorilla's mux package, requests can be matched based on URL host, path, path prefix, schemes, header and query values, and HTTP methods.

go get github.com/gorilla/mux

Mongo DB :
----------
docker pull mongo <br/>
docker image inspect mongo  <br/>
docker run -d  --name mongo-on-docker  -p 27888:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo  <br/>

#created folder to use for storing MongoDB data  <br/>
sudo mkdir /var/dbdata  <br/>
docker run -it -d -v /var/dbdata:/data/db -p 27017:27017 --name mongodb mongo  <br/>
docker start mongodb  <br/>
sudo docker exec -it mongodb bash  <br/>

docker run --link mongo_db_name_container:mongo -p 8081:8081 -e ME_CONFIG_MONGODB_URL="mongodb://mongo:27017" mongo-express  <br/>
docker run --link mongodb:mongo -p 8081:8081 -e ME_CONFIG_MONGODB_URL="mongodb://mongo:27017" mongo-express  <br/>

After above DB setup, next time onwards use the following steps : <br/>
docker start mongodb  <br/>
docker run --link mongodb:mongo -p 8081:8081 -e ME_CONFIG_MONGODB_URL="mongodb://mongo:27017" mongo-express <br/>
mongo ui -- http://localhost:8081


mongodb://<username>:<password>@<host>:<port>/?authSource=admin
mongodb://mongoadmin:secret@localhost:27888/?authSource=admin


https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial