/*
function fnWrap() {
    function fnTouzi() {
    alert('亲，请关注我们的新产品！');
    };

    fnTouzi();
};

fnWrap();
*/

// 上面的写法还可能重名，可以改写成下面封闭函数的形式,
// 一个分号表示一条空的 js 语句是合法的，
;;;;;;
// 在代码前面加分号是为了防止代码在压缩成一行时，前面的代码在结束时没有加分号导致代码出错
/*
;(function() {
    function fnTouzi() {
        alert('亲，请关注我们的新产品！');
    };

    fnTouzi();
})();
*/

// 封闭函数装高手写法
// 方式一，匿名函数前加感叹好
/*
;!function(){
    function fnTouzi(){
        alert('亲，请关注我们的新产品！');
    };
    fnTouzi();
}();
*/

// 方式二：匿名函数前加 ～
;~function(){
    function fnTouzi() {
        alert('亲，请关注我们的新产品！');
    };
    fnTouzi();
}();