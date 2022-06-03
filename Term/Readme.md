# Monolithic

- Pros:

  - Simple to build and deploy
  - Easy to loggin, configuration management and performance monitoring
  - Component communicate faster.

- Cons:
  - Although we break code into more modules, the final just one thing.
  - Hard to maintain when source code too big or too complicated.(nice)
  - Difficult to deploy big application (nice)
    - When a little code come, the whole app must be tested, deployed for each update.
  - An error that make all app down.(nice)
  - Use the same tech stack for all modules. (nice)

# Microservice

- Pros
  - Deployment flexibility
  - Choosing technology of server flexibility
  - Scale app separately
  - An error don't make all app down.
- Cons
  - How servers know to communicate with each other (oke)
    - In past, use HTTP request,using gRPC.
  - Tracing the source of problem when debugging is difficult. (oke)
