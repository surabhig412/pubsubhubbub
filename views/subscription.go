package views

const NewPage = `
{{ define "view" }}
<!DOCTYPE html>
<html lang='en'>
   <head>
      <meta charset='utf-8'>
      <meta content='IE=edge' http-equiv='X-UA-Compatible'>
      <meta content='width=device-width, initial-scale=1' name='viewport'>
      <meta content='' name='description'>
      <meta content='' name='author'>
      <title>Welcome To PubSubHubBub !!</title>
      <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
      <script src="/assets/javascript/jquery-2.2.3.min.js"></script>
      <script src="/assets/javascript/pshb.js"></script>
      <script src="/assets/javascript/bootstrap.min.js"></script>
      <link rel="stylesheet" type="text/css" href="/assets/css/styles.css">
   </head>
   <body>
      <form class="new_subscription" id="new_subscription" action="/subscribe" accept-charset="UTF-8" method="post" onsubmit="subscriptionButton.disabled = true">
         <input name="utf8" type="hidden" value="&#x2713;" /> <br>
         <p align='center'>
         <h3 align='center' >Subscribe Here</h3>
         </p>
         <div class='form-group'>
            <div class='row'>
               <div class='col-xs-12' id='chk-box'>
                  <label><h3><input type="checkbox" value="1" name="subscribe_option" id="subscribe_option"/>
                  Subscribe</h3></label>
                  <br>
               </div>
            </div>
         </div>
         <div id='option_subscribe'>
           {{template "partial" .}}
         </div>
         <div class='form-group'>
            <div class='row'>
               <div class='col-xs-12' id='chk-box'>
                  <label><h3><input type="checkbox" value="1" name="unsubscribe_option" id="unsubscribe_option" unchecked/>
                  Unsubscribe</h3></label>
                  <br>
               </div>
            </div>
         </div>
         <div id='option_unsubscribe'>
           {{template "partial" .}}
         </div>
         <div class='form-group'>
            <div class='input-group'>
               <input placeholder="CALLBACK URL" class="form-control disable_copy_paste" maxlength="30" autocomplete="off" type="text" name="subscription_callback_url" id="subscription_callback_url" />
               <div class='input-group-addon'>
                  <span class='glyphicon glyphicon-user'></span>
               </div>
            </div>
         </div>
         <div class='form-group'>
            <input type="submit" name="commit"  class="btn btn-success btn-block" id="subscriptionButton" data-disable-with="Subscribing..." />
         </div>
      </form>
   </body>
</html>
{{end}}
`

const PartialTemp = `
{{ define "partial" }}
       <div>
          {{range $key, $value := .data}}
          <p><h3>{{$key}}</h3></p>
            {{range $k, $v := $value}}
              {{if eq $k "Topics"}}
                  {{range $i, $val := $v}}
                    <div class='row'>
                      <div class='col-xs-12' id='chk-box'>
                        <t><input type="checkbox" value={{$val}}_{{$key}} name="topic" id={{$val}}_{{$key}} />
                        {{$val}} </t>
                        <br>
                      </div>
                    </div>
                  {{end}}
           {{end}}
         {{end}}
  {{ end }}
 </div>
{{ end }}
`
