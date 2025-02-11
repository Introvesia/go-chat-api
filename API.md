# API References

## User authentication

### Login

```bash
curl -X POST http://localhost:8082/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin", "password":"password123"}'
```

### Register

```bash
curl -X POST http://localhost:8082/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john_doe", "email":"john-doe@example.com", "password":"password123", "first_name":"John", "last_name":"Doe"}'
```

## Channels

### List of channels

```bash
curl -X POST http://localhost:8082/channels \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json"
```

### List of channel's members

```bash
curl -X POST http://localhost:8082/channels/{id}/members \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json"
```