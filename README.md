# URLcommentary
URLcommentary is a tool designed to enable users to comment on any website, even those where the comment section is disabled. This initiative aims to promote democratic discussion and enhance community interaction across the web.

## Objectives
- **Comment on any URL independent of the webpage**
- **Bring transparency to comments and reviews**
- **Richer community interactions**

<!--
## Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/nimilgp/URLcommentary
    ```

2. **Navigate to the Directory**

    ```bash
    cd URLcommentary/extension
    ```

3. **Load the Extension in Chrome**

    - Open Chrome and navigate to `chrome://extensions/`.
    - Enable "Developer mode" in the top right corner.
    - Click "Load unpacked" and select the directory containing your extension files.

## Usage

Once the extension is loaded,  users can then post and view comments, which will be displayed in the new comment section injected into the webpage.
-->
## Database setup
1. **Install PostgreSQL**
2. **Create a user(postgres) and database(urlc)**
3. **Add database connection string in the config file .env for dsn value**

## Running API server

1. **Clone the Repository**

    ```bash
    git clone https://github.com/nimilgp/URLcommentary
    ```

2. **Navigate to the Directory**

    ```bash
    cd URLcommentary/
    ```
3. **List out Make Commands**

   ```bash
   make help
   ```
4. **Run server**
   ```bash
   make run
   ```
