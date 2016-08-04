package views

const AdminPage = `
<!DOCTYPE html>
<html lang='en'>
<head>
<div class="row">
  <div class="medium-7 large-6 columns">
    <div class="callout secondary">
      <h5>Add your Topics..</h5>
      <form method="post" action="/addTopic/">
       <input name="publisher" placeholder="Your Name" required>
        <input name="topic1" placeholder="Your First Topic" required>
        <input name="topic2" placeholder="Your Second Topic" required>
        <input name="topic3" placeholder="Your Third Topic" required>
        <input type="submit" value="Add" class="button">
      </form>
    </div>
    <hr/>
    <h5>List of topics</h5>
    <div >
          {{range $i, $elem := .Topics}}

            <div class="callout">
              <p>{{$elem}}</p>
              <a class="close-button" aria-label="Dismiss alert" href="/delete/{{$i}}">
                <span aria-hidden="false"><i class="fa fa-trash"></i></span>
              </a>
            </div>

          {{end}}
    </div>


  </div>

</div>
</form>
</body>
</html>
`
