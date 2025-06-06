otel-cli span --service demo --name "demo-span" \
         --endpoint localhost:4317 \
         --no-tls-verify --protocol grpc  --status-code "ok" \
         --status-description "Ca marche !!" --verbose