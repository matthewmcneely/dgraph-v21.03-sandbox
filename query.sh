#!/bin/sh

read -d '' QUERY << EOM
queryProduct(\$filter: CategoryFilter, \$order: ProductOrder) {
    queryProduct(order: \$order) @cascade(fields: [\"category\"]) {
      id
      currentPrice
      category(filter: \$filter) {
        name
      }
    }
  }
EOM

read -d '' VARIABLES << EOM
{
    "filter": {
      "name": {
        "eq": "Tops"
      }
    },
    "order": {
      "asc": "currentPrice"
    }
  }
EOM

generate_post_data()
{
  cat <<EOF
  {"query": "$QUERY","variables": $VARIABLES}
EOF
}

echo $(generate_post_data)


curl -v --data "$(generate_post_data)" -H "Content-Type: application/json" -X POST localhost:8080/graphql