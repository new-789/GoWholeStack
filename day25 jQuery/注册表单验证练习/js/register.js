$(function(){
    var error_name = true;
    var error_pwd = true;
    var error_check_pwd = true;
    var error_email = true;
    var error_check = true;

    var $name = $('#user_name');
    var $pwd = $('#pwd');
    var $cpwd = $('#cpwd');
    var $email = $('#email');
    var $allow = $('#allow');

    $name.blur(function() {
        // 该函数用来验证输入内容的正确性
        check_user_name();
    }).click(function(){
        // 隐藏输入框下的错误提示信息
        $(this).next().hide();
    });

    $pwd.blur(function() {
        check_pwd();
    }).click(function(){
        $(this).next().hide();
    });

    $cpwd.blur(function(){
        check_cpwd();
    }).click(function(){
        $(this).next().hide();
    });

    $email.blur(function(){
        check_email();
    }).click(function(){
        $(this).next().hide();
    });

    $allow.click(function(){
        // 获取 checked 多选框属性值方法,勾选返回 true 否则 false
        if($(this).prop('checked')){
            error_check = false;
            $(this).siblings('p').hide();
        }else{
            error_check = true;
            $(this).siblings('p').html('请勾选同意,否则您将无法注册').show();
        };
    });

    function check_user_name(){
        var reg = /^\w{6,15}/;
        var val = $name.val();

        if (val ==''){
            $name.next().html('用户名不能为空').show();
            error_name = true;
            return;
        };
        if (reg.test(val)){
            $name.next().hide();
            error_name = false;
        }else{
            $name.next().html('用户名是6到15个英文或数字，还可包含“_”').show();
            error_name = true;
        };
    };

    function check_pwd(){
        var reg = /^[\w!@#$%^&*]{6,15}$/;
        var val = $pwd.val();

        if (val == ''){
            $pwd.next().html('密码不能为空').show();
            return;
        }
        if(reg.test(val)){
            $pwd.next().hide();
            error_pwd = false;
        }else {
            $pwd.next().html('密码由6到15位字母、数字和特殊字符组成').show();
            error_pwd = true;
        };
    };

    function check_cpwd(){
        // 定义正则规则
        var reg = /^[\w!@#$%^&*]{6,15}/;
        var val = $cpwd.val();

        if (val == ''){
            $cpwd.next().html('确认密码不能为空').show();
            return;
        };
        if (reg.test(val)){
            $cpwd.next().hide;
            error_check_pwd = false;
        }else{
            $cpwd.next().html('密码由6到15位字母、数字和特殊字符组成').show();
            error_check_pwd = true;
        };
    };

    function check_email(){
        var reg = /^[a-z0-9][\w\.\-]*@[a-z0-9\-]+(\.[a-z]{2,5}){1,2}$/i;
        var val = $email.val();

        if (val == ''){
            $email.next().html('邮箱不能为空').show();
            return;
        };
        if (reg.test(val)){
            $email.next().hide();
            error_email = false;
        }else{
            $email.next().html('您输入的邮箱不合法，请重新输入').show();
            error_email = true;
        };
    };

    // 阻止鼠标点击事件
    $('.reg_form form').submit(function(){
        // 如果都填写正确则可以正常注册，否则无法提交注册信息，否则不可提交
        if (error_name==false && error_pwd == false && error_check_pwd == false && error_email == false && error_check == false)
        {
            return true;
        }else{
            return false;
        }
    })
});