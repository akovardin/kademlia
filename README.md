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

![](https://photos-6.dropbox.com/t/2/AADnMbV-_xbbiQhAZlYENq27kW796UK7J_3YFvo-L0JPeA/12/750049/png/32x32/3/1499641200/0/2/kademlia.png/ENaoXhjpoLvbBCACKAI/YXczWDu4JjtF-sgMkQUhDQfIG8Xsjej7tvReIDZC_Wc?dl=0&size=1600x1200&size_mode=3)
