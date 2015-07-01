function $(selector) {
    return document.querySelector(selector);
}

function $$(selector) {
    return document.querySelectorAll(selector);
}

function cancelEditor() {
    $("#edit_btn").style.display = "inline-block";
    $("#save_btn").style.display = "none";
    $("#cancel_btn").style.display = "none";
    $(".editor_buttons").style.display = "none";

    var element = $("[data-editable=true]");

    var editable_elements = $$(".editable");
    for(var i=0; i<editable_elements.length; i++) {
        if(editable_elements[i].outerHTML.indexOf("<div")===0)
            editable_elements[i].removeEventListener("keydown", editorKeyPress);
        else if(editable_elements[i].outerHTML.indexOf("<pre")===0)
            editable_elements[i].removeEventListener("keydown", codeEditorKeyPress);
        editable_elements[i].contentEditable = false;
    }

    element.innerHTML = element.innerHTML.replace(/class="editable" contenteditable="false"/g, '');
    element.innerHTML = element.innerHTML.replace(/<div>([\s\S]*?)<\/div>/img, "$1");
}

window.onload = function() {
    $("#edit_btn").addEventListener("click", function(){
        $("#edit_btn").style.display = "none";
        $("#save_btn").style.display = "inline-block";
        $("#cancel_btn").style.display = "inline-block";

        $(".editor_buttons").style.display = "inline-block";

        $(".editor_buttons .btn.undo").addEventListener("click", function(){
            document.execCommand("undo", false, null);
        })

        $(".editor_buttons .btn.redo").addEventListener("click", function(){
            document.execCommand("redo", false, null);
        })

        $(".editor_buttons .btn.bold").addEventListener("click", function(){
            document.execCommand("bold", false, null);
        })

        $(".editor_buttons .btn.italic").addEventListener("click", function(){
            document.execCommand("italic", false, null);
        })

        $(".editor_buttons .btn.underline").addEventListener("click", function(){
            document.execCommand("underline", false, null);
        })

        $(".editor_buttons .btn.h1").addEventListener("click", function(){
            document.execCommand("formatBlock", false, "h1");
        })

        $(".editor_buttons .btn.h2").addEventListener("click", function(){
             document.execCommand("formatBlock", false, "h2");
        })

        $(".editor_buttons .btn.h3").addEventListener("click", function(){
             document.execCommand("formatBlock", false, "h3");
        })

        $(".editor_buttons .btn.h4").addEventListener("click", function(){
            document.execCommand("formatBlock", false, "h4");
        })

        $(".editor_buttons .btn.clear").addEventListener("click", function(){
            document.execCommand("formatBlock", false, "p");
        })

        $(".editor_buttons .btn.ol").addEventListener("click", function(){
            document.execCommand("insertOrderedList", false, false);
        })

        $(".editor_buttons .btn.ul").addEventListener("click", function(){
            document.execCommand("insertUnorderedList", false, false);
        })

        $(".editor_buttons .btn.code_block").addEventListener("click", function(){
            document.execCommand("insertHTML", false, "</div><pre></pre><div>");
        })


        var element = $("[data-editable=true]");
        element.innerHTML =  "<div class=\"editable\">"+element.innerHTML.replace(/<pre.*?>([\s\S]*?)<\/pre>/img, "</div><pre class=\"editable\">$1</pre><div class=\"editable\">")

        var editable_elements = $$(".editable");
        for(var i=0; i<editable_elements.length; i++) {
            if(editable_elements[i].outerHTML.indexOf("<div")===0)
                editable_elements[i].addEventListener("keydown", editorKeyPress);
            else if(editable_elements[i].outerHTML.indexOf("<pre")===0)
                editable_elements[i].addEventListener("keydown", codeEditorKeyPress);
            editable_elements[i].contentEditable = true;
        }
    });
    $("#cancel_btn").addEventListener("click", function(){
        while(document.execCommand("undo")){}
        cancelEditor();
    });

    $("#save_btn").addEventListener("click", function(){
        cancelEditor();

        var content = document.querySelector("[data-editable=true]").innerHTML;

        console.log(content);

        var data = new FormData();
        data.append("content", content);

        var xhr = new XMLHttpRequest();

        xhr.open("POST", "/api/post/", false);
//        xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
        xhr.send(data);
    });
}

function editorKeyPress(e) {

    if(e.keyCode==13)
    {
        e.returnValue = false;
        document.execCommand("insertHTML", false, "\n");
    }
    if(e.keyCode==9) { // tab
        e.returnValue = false;
        document.execCommand("insertHTML", false, "&emsp;");
    }
    console.log("editor key pressed!");
}

function codeEditorKeyPress(e) {
    var selection = window.getSelection();
    var selection2 = selection;
    if(e.keyCode==9) { // tab
        e.returnValue = false;
        if(selection.toString().length>0)
        {
            console.log("tab with selection");
            document.execCommand("insertHTML", false, selection.toString().replace(/^([\s\S]*?)$/gm, "\t$1"));
        } else
            document.execCommand("insertHTML", false, "\t");
    }
    console.log(selection, selection2);
    console.log("code editor key pressed!");
}