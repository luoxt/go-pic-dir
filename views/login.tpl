<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:th="http://www.thymeleaf.org">

<head>
    <title>田田云客户登录系统</title>
    <meta charset="UTF-8" />
    <!--  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="renderer" content="webkit" />
    <meta http-equiv="Cache-Control" content="no-siteapp" /> -->
    <link href="/static/css/login.css?v=20181112" rel="stylesheet" />
</head>

<body>
    <div class="login-header">
        <div class="header-con">
            <i class="iconfont icon-logo"></i>
            <div class="con-line"></div>
            <span class="con-name">田田云客户登录系统</span>
        </div>
    </div>
    <div class="main-content">
        <div class="header">田田云客户登录系统</div>
        <form class="login-form" action="/account" method="post">
            <input type="hidden" name="redirect" value="{{.redirect}}" />
            <div class="item-box" id="itemBox">
                   <div class="item">
                    <i class="iconfont icon-yonghumin"></i>
                    <input name="login_account" type="text" placeholder="请填写手机号码" autocomplete="off" value="{{.login_account}}" />
                </div>
                <span class="placeholder_copy placeholder_un">请填写手机号码</span>
                <div class="item b0">
                    <i class="iconfont icon-mima"></i>
                    <input name="login_password" type="password" placeholder="请填写密码" autocomplete="off" value="{{.login_password}}" />
                </div>
                <span class="placeholder_copy placeholder_pwd">请填写密码</span>
            </div>
            <div class="login_btn_panel">
                <button class="login-btn" type="submit">
                    <span class="in"><i class="icon-loading"></i>登 录 中 ...</span>
                    <span class="on">登 录</span>
                </button>
                <div class="check-tips">
                    <p>{{.msg}}</p>
                </div>
            </div>
        </form>
    </div>
    <div class="login_footer">
        <p>© 2018 深圳田田云网络科技有限公司 版权所有</p>
    </div>

    <script src="/static/js/lib/jquery-3.2.1.min.js"></script>
    <script src="/static/js/login.js"></script>
</body>

</html>