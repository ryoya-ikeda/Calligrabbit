function jump1() {
    location.href = "index3.html";
}

function imgchange() {
    setTimeout("jump1()", 333);
    document.img01.src = "";
}

//--></script>


<P><A href="JavaScript:imgchange()">
    <IMG src="　静止画像のリンク先　" width="xxx" height="xxx" border="0" name="img01" /></A></P>