document.write('<script type="text/javascript" src="../../bootstrap/js/authorization/ksort.js"></script>');
document.write('<script type="text/javascript" src="../../bootstrap/js/authorization/crypto-js.min.js"></script>');
document.write('<script type="text/javascript" src="../../bootstrap/js/authorization/hmac-sha256.js"></script>');
document.write('<script type="text/javascript" src="../../bootstrap/js/authorization/enc-base64.min.js"></script>');
document.write('<script type="text/javascript" src="../../bootstrap/js/jquery.cookie.min.js"></script>');


function GenerateAuthorization(path, method, params) {
    let key = "admin";
    let secret = "12878dd962115106db6d";

    let date = new Date();
    let datetime = date.getFullYear() + "-" // "Year"
        + ((date.getMonth() + 1) > 10 ? (date.getMonth() + 1) : "0" + (date.getMonth() + 1)) + "-" // "Month"
        + (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " " // "Day"
        + (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) + ":" // "Hour"
        + (date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes()) + ":" // "Minute"
        + (date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds()); // "Second"

    let sortParamsEncode = decodeURIComponent(jQuery.param(ksort(params)));
    let encryptStr = path + "|" + method.toUpperCase() + "|" + sortParamsEncode + "|" + datetime;
    let digest = CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA256(encryptStr, secret));
    return {authorization: key + " " + digest, date: datetime};
}

function IsJson(str) {
    if (typeof str == 'string') {
        try {
            let obj = JSON.parse(str);
            if (typeof obj == 'object' && obj) {
                return true;
            } else {
                return false;
            }
        } catch (e) {
            console.log('error：' + str + '!!!' + e);
            return false;
        }
    }
    console.log('It is not a string!')
}

function AjaxError(response) {
    let errCode = response.status;
    let errMsg = response.responseText;

    if (errCode === 401) { // Jump to login page
        // Close the current interface
        parent.window.close();
        window.open("/login");
        return;
    }

    if (IsJson(response.responseText)) {
        const errInfo = JSON.parse(response.responseText);
        errCode = errInfo.code;
        errMsg = errInfo.message;
    }

    $.alert({
        title: 'Error message',
        icon: 'mdi mdi-alert',
        type: 'red',
        content: 'Error code:' + errCode +'<br/>' +'Error message:' + errMsg,
    });
}

function AjaxForm(method, url, params, beforeSendFunction, successFunction, errorFunction) {
    let authorizationData = GenerateAuthorization(url, method, params);

    $.ajax({
        url: url,
        type: method,
        data: params,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
            'Authorization': authorizationData.authorization,
            'Authorization-Date': authorizationData.date,
            'Token': $.cookie("_login_token_"),
        },
        beforeSend: beforeSendFunction,
        success: successFunction,
        error: errorFunction,
    });
}

function AjaxFormNoAsync(method, url, params, beforeSendFunction, successFunction, errorFunction) {
    let authorizationData = GenerateAuthorization(url, method, params);

    $.ajax({
        url: url,
        type: method,
        data: params,
        async: false,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
            'Authorization': authorizationData.authorization,
            'Authorization-Date': authorizationData.date,
            'Token': $.cookie("_login_token_"),
        },
        beforeSend: beforeSendFunction,
        success: successFunction,
        error: errorFunction,
    });
}

function AjaxPostJson(url, params, beforeSendFunction, successFunction, errorFunction) {
    let authorizationData = GenerateAuthorization(url, "POST", params);

    $.ajax({
        url: url,
        type: "POST",
        data: JSON.stringify(params),
        headers: {
            'Content-Type': 'application/json; charset=utf-8',
            'Authorization': authorizationData.authorization,
            'Authorization-Date': authorizationData.date,
            'Token': $.cookie("_login_token_"),
        },
        beforeSend: beforeSendFunction,
        success: successFunction,
        error: errorFunction,
    });
}
