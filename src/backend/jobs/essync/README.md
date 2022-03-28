### 创建字段映射

```shell
curl -XPOST http://127.0.0.1:9200/canal/_mapping -H 'Content-Type:application/json' -d'
{
    "properties": {
        "id": {
            "type": "long"
        },
        "is_state": {
            "type": "long"
        },
        "is_del": {
            "type": "long"
        },
        "utime": {
            "type": "date",
            "format": "strict_date_optional_time||epoch_millis"
        },
        "ctime": {
            "type": "date",
            "format": "strict_date_optional_time||epoch_millis"
        },
        "title": {
            "type": "text",
            "analyzer": "ik_max_word",
            "search_analyzer": "ik_smart"
        },
        "intro": {
            "type": "text",
            "analyzer": "ik_max_word",
            "search_analyzer": "ik_smart"
        },
        "data": {
            "type": "text",
            "analyzer": "ik_max_word",
            "search_analyzer": "ik_smart"
        }
    }

}'
```
