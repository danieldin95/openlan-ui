var express = require('express');
var router = express.Router();

router.get('/', (req, res) => {
  res.send(JSON.stringify({
    "type": "force",
    "categories": [
      {
        "name": "Virtual Switch"
      },
      {
        "name": "Access Point"
      }
    ],
    "nodes": [
      {
        "name": "YunStack",
        "value": 1,
        "symbolSize": 20,
        "category": 0,
        "id": 0,
        "label": {"show":  true}
      },
      {
        "name": "HongKong",
        "value": 2,
        "symbolSize": 20,
        "category": 0,
        "id": 1,
        "label": {"show":  true}
      },
      {
        "name": "Singapore",
        "value": 20,
        "symbolSize": 20,
        "category": 0,
        "id": 2,
        "label": {"show":  true}
      },
      {
        "name": "YunDev-001",
        "value": 10,
        "symbolSize": 10,
        "category": 1,
        "id": 3
      }
    ],
    "links": [
      {
        "source": 0,
        "target": 1,
        "widget": 10
      },
      {
        "source": 1,
        "target": 2
      },
      {
        "source": 3,
        "target": 0
      },
      {
        "source": 1,
        "target": 0
      },
      {
        "source": 2,
        "target": 1
      }
    ]
  }));
});

module.exports = router;