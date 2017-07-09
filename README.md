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

![](https://photos-5.dropbox.com/t/2/AACTI4dUbVns06VsNu0bW60RW1-npnVRlYfaOm0nNR18Kg/12/750049/png/32x32/3/1499641200/0/2/kademlia.png/ENaoXhjnoLvbBCACKAI/cJBB_ujbysab_RbYtbjY8ciRXDkjndJXcFsdr_Q1Nek?dl=0&size=1600x1200&size_mode=3)
