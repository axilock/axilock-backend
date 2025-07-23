# Axilock Backend

Axilock Backend is a backend service for Axilock push protection application which prevents secrets from leaking to git repositories.

## Setup Instructions

For detailed setup instructions, please refer to the [self-host documentation](https://docs.axilock.ai/secret-prevention/self-host/backend/).


### 2.1 Configure Github App

You will need to create a github app to authenticate with the github api. You can do this by following the instructions [here](https://docs.github.com/en/apps/creating-github-apps/about-creating-github-apps/about-creating-github-apps). 
This github app will be used to authenticate with the github api to get the commit metadata which will give the coverage of how haw users are actively using the secrets protection. You need to get the values of `GITHUB_APP_ID`, `GITHUB_CLIENT_SECRET` and `GITHUB_CLIENT_ID` from the github app settings.

!!! Info
    You will also need the private key file for the github app to be stored in the root directory of the project and name the file as `axilock.pem` which is the default name used by the backend.  

### 2.2 Configure Environment Variables 

Fill the .env file with your specific configuration:

```env
HTTP_SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:8090
RUNNING_ENV=<dev/prod>
GITHUB_APP_ID=<Enter your github app id>
GITHUB_CLIENT_SECRET=<Enter your github client secret>
GITHUB_CLIENT_ID=<Enter your github client id>
```

### 3. Start Services with Docker Compose

The application uses Docker Compose to manage services. Start all services with:

```bash
docker-compose up -d
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


## Enterprise Support

For enterprise support and for custom secrets support, schedule a call with us [Schedule a call](https://cal.com/axilock/support?overlayCalendar=true).

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## Contact

For support or questions, please open an issue in the [GitHub repository](https://github.com/axilock/axilock-backend).
