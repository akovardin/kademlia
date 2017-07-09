#Kademlia

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

![](https://dl.dropboxusercontent.com/1/view/06isycp8r0y5z73/Apps/Shutter/%C3%90%C2%92%C3%91%C2%8B%C3%90%C2%B4%C3%90%C2%B5%C3%90%C2%BB%C3%90%C2%B5%C3%90%C2%BD%C3%90%C2%B8%C3%90%C2%B5_014.png)