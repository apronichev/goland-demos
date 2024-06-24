## Using the HTTP Client

1. **Code Highlighting**
    - Open `requests.http` and `requests.rest` in IntelliJ IDEA. The syntax will be highlighted automatically.

2. **Code Completion**
    - With `openapi.yaml` in your project, IntelliJ IDEA will provide code completion for hosts, method types, header fields, and endpoints. Start typing in the HTTP request files and use `Ctrl+Space` to trigger code completion.

3. **Code Folding**
    - You can fold sections of your HTTP request files by clicking the icon in the gutter next to the request parts. This works for entire requests, headers, and response handler scripts.

4. **Reformat Requests**
    - Use `Ctrl+Alt+L` (or `Cmd+Alt+L` on macOS) to reformat the HTTP requests according to your code style.

5. **Inline Documentation**
    - Hover over request header fields or OpenAPI doc tags to see inline documentation and descriptions.

6. **Viewing Structure**
    - Open the structure view (`Alt+7` or `Cmd+7` on macOS) to see the structure of your HTTP request files, including requests and their parts.

7. **Language Injections**
    - IntelliJ IDEA supports language injections in Web languages inside the request message body. For example, the JSON body in the POST request will have proper syntax highlighting and validation.

8. **Live Templates**
    - Use live templates to quickly generate request snippets. Type the template abbreviation and press `Tab` to expand it. For example, type `get` and press `Tab` to insert a GET request template.


### Making live template

Live templates in IntelliJ IDEA can significantly speed up the process of writing HTTP requests. Below is an example of how to create a live template for an HTTP GET request in IntelliJ IDEA. This template will include placeholders for the HTTP method, URL, headers, and query parameters.

### Step-by-Step Guide to Create a Live Template for HTTP Request

1. **Open Live Template Settings**
   - In IntelliJ IDEA, go to **File > Settings** (or **IntelliJ IDEA > Preferences** on macOS).
   - Navigate to **Editor > Live Templates**.

2. **Create a New Template Group**
   - Click the `+` icon to add a new group.
   - Name the group `HTTP`.

3. **Create a New Live Template**
   - Select the `HTTP` group and click the `+` icon to add a new live template.
   - Enter the template abbreviation, description, and template text as follows:

   **Abbreviation**: `httpget`

   **Description**: `Template for HTTP GET request`

   **Template Text**:
   ```http
   GET $URL$ HTTP/1.1
   Host: $HOST$
   Accept: application/json
   $END$
   ```

4. **Define Variables**
   - Click the **Edit variables** button to define the variables used in the template.
   - For `URL`, set the expression to `complete()`.
   - For `HOST`, set the expression to `complete()`.
   - Leave `END` as is.

5. **Set Context**
   - Click **Define** and select `HTTP Request`. This will ensure the template is available when editing HTTP request files.

6. **Save the Template**
   - Click **OK** to save the template.

### Using the Live Template

1. **Create or Open an HTTP Request File**
   - Create a new `.http` or `.rest` file or open an existing one in your project.

2. **Invoke the Live Template**
   - Type the abbreviation `httpget` and press `Tab` or `Enter`. The template will expand, and you can fill in the placeholders.

### Example of Using the Template

After invoking the `httpget` template, the expanded template will look like this:

```http
GET http://localhost:8080/users HTTP/1.1
Host: localhost:8080
Accept: application/json
```

You can then customize the URL, Host, and add any additional headers or query parameters as needed.

### Additional Templates

You can create additional live templates for other HTTP methods (POST, PUT, DELETE) using a similar approach. Here’s an example for a POST request:

**Abbreviation**: `httppost`

**Description**: `Template for HTTP POST request`

**Template Text**:
```http
POST $URL$ HTTP/1.1
Host: $HOST$
Content-Type: application/json
Accept: application/json

{
  $BODY$
}
```

**Variables**:
- `URL`: `complete()`
- `HOST`: `complete()`
- `BODY`: Leave empty
- `END`: Leave as is

**Context**:
- Set to `HTTP Request`.
