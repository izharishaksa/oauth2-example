# Client-Server Interaction for OAuth 2.0 Authentication

This project consists of a client and a server. The client initiates the OAuth 2.0 authentication process by redirecting
users to the server for authentication and authorization. Once the user grants permission, the server redirects back to
the client with an authorization code, which the client exchanges for an access token.

## Client

1. Open the client directory in your terminal:

```bash
cd client
```

2. Open the `main.go` file in your preferred code editor and update the following variables:

- `clientId`: Your OAuth 2.0 client ID.
- `clientSecret`: Your OAuth 2.0 client secret.
- `callbackUrl`: The callback URL where the authorization server redirects the user after authentication. Make sure this
  URL is registered with your OAuth 2.0 server.
- `tokenUrl`: The URL of the OAuth 2.0 token endpoint where the authorization code will be exchanged for an access
  token.

3. Save your changes.

4. Start the client application by running the following command:

```bash
go run main.go
```

5. Access the client application in your web browser by navigating to http://localhost:9191. You will see a link to "
   Sign Up Using Kancyl."

6. Click the "Sign Up Using Kancyl" link, which will initiate the OAuth 2.0 authorization process. You will be
   redirected to the authorization server for authentication and authorization.

## Server

1. Open the server directory in a new terminal window:

```bash
cd server
```

2. Open the main.go file in your preferred code editor and configure the server to handle OAuth 2.0 authentication and
   authorization. Implement the authorization and token endpoints on the server to facilitate the OAuth 2.0 flow.

3. Start the server by running the following command:

```bash
go run main.go
```

4. The server should be running and listening for incoming OAuth 2.0 authorization requests from the client.

## OAuth 2.0 Flow

1. The client application (client) initiates the OAuth 2.0 flow by sending the user to the authorization server (server)
   for authentication and authorization.

2. The authorization server handles user authentication and authorization.

3. After successful authentication and authorization, the authorization server redirects the user back to the client
   application's callback URL (callbackUrl) with an authorization code.

4. The client application receives the authorization code and exchanges it for an access token by sending a request to
   the OAuth 2.0 token endpoint (tokenUrl) on the authorization server.

5. The authorization server validates the authorization code and responds with an access token.

6. The client application can now use the obtained access token to make authorized API requests to protected resources
   on the server.

This project demonstrates a basic OAuth 2.0 client-server interaction for user authentication and authorization. In a
production environment, you should handle errors, implement security best practices, and further customize the client
and server applications as needed.