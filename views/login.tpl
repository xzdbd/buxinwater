<!DOCTYPE html>

<html>
<head>
  <title>Login - Box Portal</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
  <link href="//cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.2.10/semantic.css" rel="stylesheet">
  
  <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.2.10/semantic.js"></script>
  <style type="text/css">
  #header-image {
    background: url(http://www.esri.com/~/media/Images/Content/landing-pages/arcgis-online/arcGIS/main-banner.jpg)  no-repeat;
      height: 260px
    }

    #footer{
      text-align: center;
    }
    body {
      background-color: #364150;
    }
    body > .grid {
      height: 100%;
    }
    .image {
      margin-top: -100px;
    }
    .column {
      max-width: 450px;
    }
  </style>
  <script>
  $(document)
    .ready(function() {
      $('.ui.form')
        .form({
          fields: {
            username: {
              identifier  : 'username',
              rules: [
                {
                  type   : 'empty',
                  prompt : 'Please enter your username'
                } 
              ]
            },
            password: {
              identifier  : 'password',
              rules: [
                {
                  type   : 'empty',
                  prompt : 'Please enter your password'
                }
              ]
            }
          }
        })
      ;
    })
  ;
  </script>
</head>
<body>
  <div class="ui middle aligned center aligned grid">
    <div class="column">
      <h2 class="ui teal image header">
        <div class="content">
          登入系统
        </div>
      </h2>
      <form class="ui large form" method="post" enctype="application/x-www-form-urlencoded">
        <div class="ui stacked segment">
          <div class="field">
            <div class="ui left icon input">
              <i class="user icon"></i>
              <input type="text" name="username" placeholder="Username">
            </div>
          </div>
          <div class="field">
            <div class="ui left icon input">
              <i class="lock icon"></i>
              <input type="password" name="password" placeholder="Password">
            </div>
          </div>
          <div class="ui fluid large teal submit button">登陆</div>
        </div>
        <div class="ui error message" style="{{if .isLoginFail}}display: block{{else}}{{end}}">
          {{if .isLoginFail}}
          <ul class="list">
            <li>
              Your username and password did not match. Please try again.
            </li>
          </ul>
          {{else}}
          {{end}}
        </div>
      </form>
      <div class="ui message">
      New to us? <a href="#">Sign Up</a>
      </div>
      <div class="ui blue label">2016 © xzdbd All Rights Reserved.</div>
      <div class="ui orange label">Version: 0.0.1</div>
    </div>
  </div>
</body>
<div></div>
</html>
