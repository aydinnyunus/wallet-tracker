docker pull neo4j:4.4.2
docker-compose up -d neo4j
# Wait till neo4j is available
sleep 60

# Define unique constraints
docker exec -it neo4j cypher-shell -u neo4j -p letmein -d neo4j 'CREATE CONSTRAINT IF NOT EXISTS ON (t:Transaction) ASSERT t.hash IS UNIQUE;'
docker exec -it neo4j cypher-shell -u neo4j -p letmein -d neo4j 'CREATE CONSTRAINT IF NOT EXISTS ON (a:Address) ASSERT a.id IS UNIQUE;'
docker exec -it neo4j cypher-shell -u neo4j -p letmein -d neo4j 'CREATE INDEX FOR (t:Transaction) ON (t.timestamp);'

docker-compose stop


