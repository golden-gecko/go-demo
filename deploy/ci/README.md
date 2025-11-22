# CI

## Get certificates from server

```bash
openssl s_client -showcerts -servername registry.com.gecko -connect registry.com.gecko:443 -CAfile deploy/ci/nginx/ssl/ca-cert.pem </dev/null
```

## Show certificate subject

```bash
openssl x509 -noout -subject -in deploy/ci/nginx/ssl/ca-cert.pem
```

## Skip verification

```bash
curl -v --ssl-no-revoke --cacert deploy/ci/nginx/ssl/ca-cert.pem https://registry.com.gecko/
```
