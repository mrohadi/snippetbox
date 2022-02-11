{{template "base" .}}

{{define "title"}}Snippet #{{.Snippet.ID}}{{end}}

{{define "body"}}
<div class="snippet">
    <div class="metadata">
        <strong>{{.Snippet.Title}}</strong>
        <span>#{{.Snippet.ID}}</span>
    </div>
    <pre><code>{{.Snippet.Content}}</code></pre>
    <div class="metadata">
        <time>Created: {{.Snippet.Created}}</time>
        <time>Expires: {{.Snippet.Expires}}</time>
</div>
{{end}}