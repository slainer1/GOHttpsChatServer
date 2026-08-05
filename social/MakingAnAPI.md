Making a RESTful API
1. To start making a RESTful API, start with your goals and your structure. You must plan meticulously about how the structure of your API is going to be.
A smart structure separates/nests each essence of a RESTful API to be structured. 
    --Example: have an API folder that holds routes and implementations, have
a migrations folder for pure migrations in the database, scripts folder to make your life easier, and modded/configuration folders (or in the root) to customize libraries

2. As much as structure matters, documentation also matters. Each heandler/method (not helper methods) must be thoroughly documented and easily accessible through curl commands.
It is absolutely crucial in a team environment to know what methods and handlers route to. Having an easily accessible UI like Swagger makes this process
much easier to debug errors.

3. Speaking of errors, you must account for errors in each method and write 2-3 edge cases. Each function should throw an error in different scenarios.


Making Routes and methods
1. First start off with your documentation. Define each element you are going to route in the documentation first when working with GOLang.
2. Write your route in the mounting method of the API. 
3. Route that method through a clean NEW program page so each element/feature stays clean
4. The language of RESTApis is usually JSON, so there is a lot of writing and reading JSON from said features
5. If needed to create a local storage to SQL inject into the database, or to hold the data there locally
6. CI/CD this stuff

Commands of REST HTTP verbs:

1. GET / https://example.com → Retrieves specific user 12.
2. POST / https://example.com → Creates a new user.
3. PUT / https://example.com → Replaces user 12 entirely.
4. PATCH / https://example.com → Partially updates user 12.
5. DELETE / https://example.com → Deletes user 12

TIPS and Rules to follow
1. Robust security and AUTH. Encrypt all data in transit and use JWT, OAuth2, rate limit, and CORS
2. Have a predictable URL structure so users can easily navigate the API: nouns not verbs, logical nesting, idempotency
3. Clear data handling and versioning, include the version in the URL path to protect production integrations, picj one camelCase or snake_case, pgination, filtering/sorting
4. Sementic status codes, 401,500,109 etc. You need to know which errors are which
5. Detailed Error payloads (a block of JSON) always return a structured JSON block explaning the failure, not raw backend stack traces
6. Structured logs and documentation (swagger)

Structured logging
1. {Key-value pair: error "," properties: ","}

Vocabulary
Payload: Typically a block of JSON, built from a class or a struct that defines what data is going to be transferred
HTTPS methods: POST, PUT, PATCH, GET, DELETE, the language used to deliver payloads
BCRYPT: The industry standard for hashing passwords. 