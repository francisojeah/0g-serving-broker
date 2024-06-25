# 0G Serving Agent

## Setup

1. Start MySQL

   ```sh
   docker run --name mysql-8.0 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<password> -d mysql:8.0
   ```

2. Create Database `serving`

   ```sh
   docker exec -i mysql-8.0 mysql -uroot -p<password> -e "CREATE DATABASE IF NOT EXISTS serving CHARACTER SET utf8mb4;"
   ```

3. Copy and Modify Configuration File

   Copy and modify the configuration file from the [example](config-example.yaml) to suit your setup.

## Running the Agent

1. Start the Serving Agent

   Use the following command to start the serving agent:

   ```sh
   PORT=<PORT> CONFIG_FILE=<path_to_config> go run main.go
   ```

## Basic Usage Process

1. Provider Starts a Chat Bot Service

   For example, a service running at: `http://localhost:3000`.

2. Provider Registers the Service with the Agent

   ```sh
   curl -X POST http://<agent_url>/v1/provider/service -H "Content-Type: application/json" -d '{"URL": "http://localhost:3000", "inputPrice": 1, "outputPrice": 2, "Type": "HTTP", "Name": "<service_name>"}'
   ```

3. User Creates an Account

   `User` creates an account to call the services registered by the `Provider`.

   ```sh
   curl -X POST http://<agent_url>/v1/user/account -H "Content-Type: application/json" -d '{"user": "<user_address>", "provider": "<provider_address>", "balance": "10000"}'
   ```

4. User Calls a Provider's Service

   For example, calling a chat bot service created by `<provider_address>` with the name `<service_name>`.

   ```sh
   curl -X POST http://<agent_url>/v1/user/retrieval/<provider_address>/<service_name> -H "Content-Type: application/json" -d '{"message": "Who won the world series in 2021?"}'
   ```

5. Provider Settles the Fee

   ```sh
   curl -X POST http://<agent_url>/v1/provider/settle
   ```

6. Provider Deletes the Service

   ```sh
   curl -X DELETE http://<agent_url>/v1/provider/service/<service_name>
   ```

7. User Checks Remaining Balance

   ```sh
   curl -X GET http://<agent_url>/v1/user/account
   ```
