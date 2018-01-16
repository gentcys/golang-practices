{{define "cookies"}}
<!DOCTYPE html>
<html>
    <head>
        <title>Show Cookies</title>
    </head>
    <body>
    {{.}}
        <h1>Cookie</h1>
        <p>Name: {{.Name}}</p>
        <p>Value: {{.Value}}</p>
        <p>Expires: {{.Expires}}</p>
    </body>
</html>
{{end}}