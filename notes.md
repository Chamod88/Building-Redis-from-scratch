### Redis Docker Container Setup

We have successfully started a Redis container on your system:
* **Image**: `redis:alpine`
* **Container Name**: `redis-server`
* **Port Mapping**: `6379:6379`

### Useful Docker Commands

1. **Check if the Redis container is running**:
   ```powershell
   docker ps
   ```

2. **Test connection with PING**:
   ```powershell
   docker exec redis-server redis-cli ping
   ```
   *Should return `PONG`.*

3. **Open an interactive Redis CLI session inside the container**:
   ```powershell
   docker exec -it redis-server redis-cli
   ```

4. **Stop the container (free up port 6379 for your Go server)**:
   ```powershell
   docker stop redis-server
   ```

5. **Start the container again**:
   ```powershell
   docker start redis-server
   ```

### Next Steps for the Go Clone
When you run your own Go Redis clone (which will listen on port `6379`), you will need to stop the Docker Redis container (`docker stop redis-server`) so that your Go application can bind to port `6379`. You can then use Docker to run `redis-cli` against your Go server:
```powershell
docker run -it --rm redis:alpine redis-cli -h host.docker.internal -p 6379
```
*Note: Make sure Go is installed on your machine to build and run the project.*