package sing

var timeLinePageTemplateHtml = `
<style>
  *
  {
    margin: 0;
    padding: 0;
  }
  .sing
  {
    background-color: #fdd;
    margin: 5px 30px;
    border: solid 4px #eaa;
    border-radius: 8px;
  }
  .sing p
  {
    display: inline-block;
  }
  .ribbon
  {
    display: inline;
    height: 100%;
    padding:5px;
    border-radius: 4px 0px 0px 4px;
    background-color: #0e0;
  }
  .singer
  {
    padding: 5px;
    width: 100px;
    background-color: #bdf;
  }

  .text
  {
    padding: 5px;
  }
   h1
  {
    margin: 30px;
  }
  form
  {
    margin: 30px;
  }
</style>
<h1>TimeLine</h1>
<form method="post" action="/post/">
  <textarea name="text"></textarea>
  <input type="submit" value="Sing">
</form>
{{range $_, $s :=  .Sings}}
  <article class="sing">
    <div class="ribbon"></div>
    <p class="singer">{{$s.Singer.Name}}</p>
    <p class="text">{{$s.Text}}</p>
  </article>
{{end}}
`
