{{template "base" .}}

{{define "title"}}Home{{end}}

{{define "body"}}
    <h2>Latest Snippets</h2>
    {{if .Snippets}}
        <table>
            <tr>
                <th>Title</td>
                <th>Creaated</td>
                <th>Id</td>
            </tr>
            {{range .Snippets}}
            <tr>
                <td><a href="/snippet?id={{.ID}}">{{.Title}}</a></td>
                <td>{{humanDate .Created}}</td>
                <td>#{{.ID}}</td>
            </tr>
            {{end}}
        </table>
    {{else}}
        <p>There is nothing to see here... yet!</p>
    {{end}}
{{end}}
