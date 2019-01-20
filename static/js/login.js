/* 登陆表单获取焦点变色 */
$(".login-form").on("focus", "input", function() {
    //$('.login_btn_panel').find(".check-tips").text('');
    $(this).closest('.item').addClass('focus');
}).on("blur", "input", function() {
    $(this).closest('.item').removeClass('focus');
});

//表单提交
$(document).ajaxStart(function() {
    $("button:submit").addClass("log-in").attr("disabled", true);
}).ajaxStop(function() {
    $("button:submit").removeClass("log-in").attr("disabled", false);
});

$("form").submit(function() {
	var self = $(this);
	self.find(".check-tips").text('');
	if(self.find("input[name=orgCode]").val()==''){
		self.find(".check-tips").text('组织编码不能为空');
		return false;
	}
	if(self.find("input[name=userName]").val()==''){
		self.find(".check-tips").text('登陆账号不能为空');
		return false;
	}
	if(self.find("input[name=passWord]").val()==''){
		self.find(".check-tips").text('登陆密码不能为空');
		return false;
	}
});

$(function() {
    //初始化选中用户名输入框
    $("#itemBox").find("input[name=userName]").focus();

    //placeholder兼容性
    //如果支持
    function isPlaceholer() {
        var input = document.createElement('input');
        return "placeholder" in input;
    }
    //如果不支持
    if (!isPlaceholer()) {
        // $()
        $(".placeholder_copy").css({
            display: 'block'
        })
        $(".placeholder_copy").click(function() {
            $(this).prev().find('input').focus()
        })
        $("#itemBox input").keydown(function() {
            $(this).parents(".item").next(".placeholder_copy").css({
                display: 'none'
            })
        })
        $("#itemBox input").blur(function() {
            if ($(this).val() == "") {
                $(this).parents(".item").next(".placeholder_copy").css({
                    display: 'block'
                })
            }
        })
    }
});