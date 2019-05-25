Demonstration of EOF responses from Vanadium RPC services.

To reproduce,

1. Create a principal:
```
go run v.io/x/ref/cmd/principal \
  create \
  --with-passphrase=false \
  --overwrite \
  /tmp/.v23 \
  hello
```

2. Start a mounttable service:
```
go run v.io/x/ref/services/mounttable/mounttabled \
  --v23.credentials /tmp/.v23 \
  --v23.tcp.address :23001
```

3. Start the server:
```
go run hello/server \
  --v23.credentials /tmp/.v23 \
  --v23.namespace.root /localhost:23001 \
  --v23.tcp.address :23000
```

4. Loop the client:
```
go build -o /tmp/client hello/client && \
for i in {1..10}; do
  /tmp/client \
    --v23.credentials /tmp/.v23 \
    --v23.namespace.root /localhost:23001 \
    --server hello_local
done
```

5. To make local changes to v.io source code for debugging, add a replace clause to go.mod.
