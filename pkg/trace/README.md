## trace

An auxiliary tool for development and debugging.

It can display real-time operation request information, running status, SQL execution, error prompts, etc. of the current page.

-`trace.go` main entry file;
-`dialog.go` processes third_party_requests records;
-`debug.go` handles debug records;

#### Data Format

##### trace_id

The ID of the current trace, for example: 938ff86be98439c6c1a7, which is convenient for searching and using.

##### request

Request information will include:

-ttl request timeout time, for example: 2s or un-limit
-method request method, for example: GET or POST
-decoded_url request address
-header request header information
-body request body information

##### response

-header response header information
-body response message
-business_code business code, for example: 10010
-business_code_msg business code information, for example: signature error
-http_code HTTP status code, for example: 200
-http_code_msg HTTP status code information, for example: OK
-cost_seconds elapsed time: unit of second, such as 0.001105661

##### third_party_requests

Each third-party http request will generate the following set of data, and multiple requests will generate multiple sets of data.

-request, the same as the above request structure
-response, same as the response structure above
-success, whether it is successful, true or false
-cost_seconds, elapsed time: unit second

Note: business_code and business_code_msg in the response are empty, because each third-party return structure is different, these two fields are empty.

##### sqls

The executed SQL information, multiple SQLs will record multiple sets of data.

-timestamp, time, format: 2006-01-02 15:04:05
-stack, file address and line number
-cost_seconds, execution time, unit: seconds
-sql, SQL statement
-rows_affected, affect the number of rows

##### debugs

-Key printed logo
-value printed value

```cassandraql
// When debugging, use this method:
p.Print("key", "value", p.WithTrace(c.Trace()))
```

Only when `p.WithTrace(c.Trace())` is added to the parameter, will it be recorded in `debugs`.

##### success

Success, true or false

```cassandraql
success = !ctx.IsAborted() && ctx.Writer.Status() == http.StatusOK
```

##### cost_seconds

Elapsed time: unit second, such as 0.001105661