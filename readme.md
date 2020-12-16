#admin portal

### how to build 
create go.mod file
update latest dependency:
    go clean --modcache
    go get github.com/tidusant/chadmin-repo@master
    go get github.com/tidusant/c3m-common@master
compile code:
    env CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o c3madmin_portal .

### run in local:
cd colis/portals/admin

env AUTH_IP=127.0.0.1 SHOP_IP=127.0.0.1 SHOP_PORT=32002 ORD_IP=127.0.0.1 ORD_PORT=32003 SESSION_URI="mongo_server_uri" SESSION_DB="mongodb_name" CHADMIN_URI="mongo_server_url" CHADMIN_DB="mongodb_name"  go run main.go 

### run in docker:
docker build -t tidusant/colis-portal-admin . && docker run -p 8081:8080 --env AUTH_IP=192.168.0.105 --name colis-portal-admin tidusant/colis-portal-admin 