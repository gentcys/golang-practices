{{define "index"}}
<!DOCTYPE html>
<html>
    <head>
        <title>Set Cookies</title>
    </head>
    <body>
        <form action="/setcookie" method="post">
            Cookie name:</br>
            <input type="text" name="name" />
            </br>
            Cookie value:</br>
            <input type="text" name="value" />
            </br>
            Expires time:</br>
            <input type="date" name="expires_date" />
            </br>
            <input type="submit" value="Submit" />
        </form>
    </body>
</html>
{{end}}