function $(selector) {
    return document.querySelector(selector);
}

function $$(selector) {
    return document.querySelectorAll(selector);
}

window.onload = function() {
    $("#edit_btn").addEventListener("click", function(){
        $("#edit_btn").style.display = "none";
        $("#save_btn").style.display = "inline-block";
        $("#cancel_btn").style.display = "inline-block";


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
        $("#edit_btn").style.display = "inline-block";
        $("#save_btn").style.display = "none";
        $("#cancel_btn").style.display = "none";

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
    });

    $("#save_btn").addEventListener("click", function(){
        $("#cancel_btn").click();
        
        console.log("ajax call..");
    });
}

function editorKeyPress(e) {

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