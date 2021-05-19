## PHP encryption algorithm corresponding to UrlSign

```php
// Sort the params key
ksort($params);

// Encode sortParams
$sortParamsEncode = http_build_query($params);

// Encrypted string rule path + method + sortParamsEncode + secret
$encryptStr = $path. $method. $sortParamsEncode. $secret

// md5 the encrypted string
$md5Str = md5($encryptStr);

// base64 encode md5Str
$tokenString = base64_encode($md5Str);

echo $tokenString;
```