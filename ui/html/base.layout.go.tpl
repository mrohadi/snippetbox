{{define "base"}}
<!doctype html>
<html lang='en'>
    <head>
        <meta charset='utf-8'>
        <title>{{template "title" .}} - Snippetbox</title>

        <!-- Link to the CSS stylesheet and favicon -->
        <link rel='stylesheet' href='/static/css/main.css'>
        <link rel='shortcut icon' href='/static/img/favicon.ico' type='image/x-icon'>

        <!-- Also link to some fonts hosted by Google -->
        <link rel="preconnect" href="https://fonts.googleapis.com">
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
        <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap" rel="stylesheet">
    </head>
    <body>
        <header>
            <h1><a href='/'>Snippetbox</a></h1>
        </header>

        <nav>
            <div>
                <a href='/'>Home</a>
                <a href='/snippet/create'>Create snippet</a>
            </div>

            <div>
                <a href='/user/signup'>Signup</a>
                <a href='/user/login'>Login</a>
                <form action='/user/logout' method='POST'>
                    <button>Logout</button>
                </form>
            </div>
        </nav>

        {{with .Flash}}
        <section>
            <div class="flash">{{.}}</div>
        </section>
        {{end}}

        <section>
            {{template "body" .}}
        </section>
        
        <!-- Invoke the footer partial template -->
        {{template "footer" .}}

    </body>
</html>
{{end}}
