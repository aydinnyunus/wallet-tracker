docker pull neo4j:4.4.2
docker compose up -d neo4j
# Wait till neo4j is available
sleep 60

# Load environment variables from .env file
if [ -f .env ]; then
  source .env
else
  echo ".env file not found"
  exit 1
fi

docker exec -e NEO4J_USERNAME="$NEO4J_USERNAME" -e NEO4J_PASSWORD="$NEO4J_PASS" neo4j cypher-shell -u "$NEO4J_USERNAME" -p "$NEO4J_PASS" "ALTER USER $NEO4J_USERNAME SET PASSWORD '$NEO4J_PASS';"

# Define unique constraints
docker exec neo4j cypher-shell -u $NEO4J_USERNAME -p $NEO4J_PASS -d neo4j 'CREATE CONSTRAINT IF NOT EXISTS ON (t:Transaction) ASSERT t.hash IS UNIQUE;'
docker exec neo4j cypher-shell -u $NEO4J_USERNAME -p $NEO4J_PASS -d neo4j 'CREATE CONSTRAINT IF NOT EXISTS ON (a:Address) ASSERT a.id IS UNIQUE;'
docker exec neo4j cypher-shell -u $NEO4J_USERNAME -p $NEO4J_PASS -d neo4j 'CREATE INDEX FOR (t:Transaction) ON (t.timestamp);'

docker compose stop


