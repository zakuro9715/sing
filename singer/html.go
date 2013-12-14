package singer

var registerPageHtml = `
<!DOCTYPE html>
<style>
  label
  {
    display: block;
    width: 100px;
  }
  form *
  {
    float: left;
  }
  form *:nth-child(odd)
  {
    clear: both;
  }
</style>
<title>Register -Sing</title>
<h1>Register</h1>
<form method="post" action="/register/">
    <label>YourName</label>
    <input type="text" name="name" required>
    <label>Introduction</label>
    <textarea name="introduction"></textarea>
  <input type="submit">
</form>
`

var viewPageTemplateHtml = `
<title>{{.Name}} -Sing</title>
  <h1>{{.Name}}</h1>
  <p>{{.Introduction}}</p>
`
