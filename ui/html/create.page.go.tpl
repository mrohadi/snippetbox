{{template "base" .}}

{{define "title"}}Create a New Snippet{{end}}

{{define "body"}}
<form action='/snippet/create' method='POST'>
    {{with .Form}}
    <div>
        <label>Title:</label>
            {{with .FormErrors.title}}
                <label class="error">{{.}}</label>
            {{end}}
        <input type='text' name='title' value="{{.Get "title"}}">
    </div>
    <div>
        <label>Content:</label>
        {{with .FormErrors.content}}
            <label class="error">{{.}}</label>
        {{end}}
        <textarea name='content'>{{.Get "content"}}</textarea>
    </div>
    <div>
        <label>Delete in:</label>
        {{with .FormErrors.expires}}
            <label class="error">{{.}}</label>
        {{end}}
        {{$exp := or (.Get "expires") "365"}}
        <input type='radio' name='expires' value='365' {{if (eq $exp "365")}}checked{{end}}> One Year
        <input type='radio' name='expires' value='7' {{if (eq $exp "7")}}checked{{end}}> One Week
        <input type='radio' name='expires' value='1' {{if (eq $exp "1")}}checked{{end}}> One Day
    </div>
    <div>
        <input type='submit' value='Publish snippet'>
    </div>
    {{end}}
</form>
{{end}}