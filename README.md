# Kademlia

Попытка реализоавать dht. 

Как запускать ноду:

```
./bin/kademlia --address="127.0.0.1:3001" --closenode="b10dc9bf98ca9648ff5fcc851692c525cc5b7aeb" --closeaddr="127.0.0.1:3000"
node id: 003cdeb33d5ec69b8038ee94aafd27d7aea4aa94
```

Как искать ноду:

```
./bin/kademlia --address="127.0.0.1:3002" --closenode="003cdeb33d5ec69b8038ee94aafd27d7aea4aa94" --closeaddr="127.0.0.1:3001" --find="b10dc9bf98ca9648ff5fcc851692c525cc5b7aeb"
result &{b10dc9bf98ca9648ff5fcc851692c525cc5b7aeb 127.0.0.1:3000}
```

Пример как работает поиск

![](kademlia.png)
