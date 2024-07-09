/**
 * NOTES:
 *  - There are still some problems related to smart contract tasks interactions that need to be solved
 *  - Client is not done yet
 *  - Dockers images of microservices still need to be tested
 *  - Here are the steps to get the app started while we keep fixing the problems
 */

 Server:
    - Each microservices will have to be started individually until container images are fixed
    - Create a .env file in eahc microservice with the proper values (you can see a template in .env.example)
    - Starts each microservices individually (wallets and tasklists will need to be start before gateway)
    - There is also to .http files in the gateway microservice to interact with backend by Rest API

 Client:
    - npm i to install dependencies
    - create .env file (you can see a template in .env.example)
    - npm run dev to start the app
