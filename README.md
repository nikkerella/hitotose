# Hitotose

```
# Mac
sudo docker cp 202408 104:/202408
sudo docker exec -i 104 /usr/bin/mongorestore --db hitotose /202408/hitotose
```

## Temp

``` mongosh
db.game.updateMany({}, { $rename: { "developer_id": "developer" } });
db.game.updateMany({}, { $rename: { "publisher_id": "publisher" } });
db.game.updateMany({}, { $rename: { "total_time": "played_time" } });
db.game.updateMany({}, { $rename: { "how_long_to_beat": "time_to_beat" } });

# db.game.updateMany({}, { $unset: { developer: "" } });
# db.game.updateMany({}, { $unset: { publisher: "" } });

db.game.find({}).forEach(function(doc) {
  if (doc.developer instanceof ObjectId) {
    db.game.updateOne(
      { _id: doc._id },
      { $set: { developer: doc.developer.toString() } }
    );
  }
});

db.game.find({}).forEach(function(doc) {
  if (doc.publisher instanceof ObjectId) {
    db.game.updateOne(
      { _id: doc._id },
      { $set: { publisher: doc.publisher.toString() } }
    );
  }
});
```
