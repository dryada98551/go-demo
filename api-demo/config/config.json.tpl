{
  "PROJECT": {
    "NAME": "blockaction",
    "JWTKEY": "jwt-secret-key"
  },
  "GIN": {
    "HOST": "0.0.0.0",
    "PORT": "8080",
    "DOMAIN": "xxx.xxx.xxx",
    "BASEPATH": "/account-api"
  },
  "DB": {
    "USER": "postgres",
    "PASSWORD": "test",
    "DBNAME": "blockaction",
    "HOST": "127.0.0.1",
    "PORT": "5431"
  },
  "REDIS": {
    "PASSWORD": "",
    "DB": 0,
    "HOST": "192.168.x.xxx",
    "PORT": "6379"
  },
  "EVENT": {
    "REDIS": {
      "PATTERN": "account"
    }
  },
  "ONEALL": {
    "PUBLICKEY": "",
    "PRIVATEKEY": "",
    "ENDPOINT": "https://app-xxxxxxxx.api.oneall.com"
  }
}