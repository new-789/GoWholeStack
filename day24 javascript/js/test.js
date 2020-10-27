window.onload=function() {
    var oBtn = document.getElementById("btn");
    var iNum1 = 12;

    oBtn.onclick=function() {
        var iNum2 = 24;
        var iRs = fnAdd(iNum1, iNum2);
        alert(iRs);
    };

    function fnAdd(a,b){
        var iRs2 = a+b;
        return iRs2;
    };
}