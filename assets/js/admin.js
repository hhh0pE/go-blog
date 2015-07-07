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

    var elements = $$("[data-editable=true][data-key]");
    for(var i=0; i<elements.length; i++)
    {
        var element = elements[i];
        element.contentEditable = false;
        element.removeEventListener("keydown", editorKeyPress);

        element.style.border = "none";
        element.style.padding = "0";

        var timer_id = parseInt(element.getAttribute('timer-id'));
        window.clearInterval(timer_id);
        element.removeAttribute('timer-id');
    }


}

var timer;

window.onload = function() {
    $("body[data-type=post] .admin-panel").innerHTML += '<div class="document_buttons">'+
            '<input type="button" id="edit_btn" value="Edit" />' +
            '<input type="button" style="display: none;" id="save_btn" value="Save" />' +
            '<input type="button" style="display: none;" id="cancel_btn" value="Cancel" />' +
        '</div>';

    $("body[data-type=post] .admin-panel").innerHTML += '<div class="editor_buttons">' +
         '<input type="button" data-command="undo" class="btn" value="Undo" />' +
        '<input type="button" data-command="redo" class="btn" value="Redo" />' +
        '&#160;&#160;&#160;&#160;' +
        '<input type="button" data-command="bold" class="btn" value="B" />' +
        '<input type="button" data-command="italic" class="btn" value="I" />' +
        '<input type="button" data-command="underline" class="btn" value="U" />' +
        '&#160;&#160;' +
        '<input type="button" data-command="formatBlock" data-argument="h1" class="btn" value="H1" />' +
        '<input type="button" data-command="formatBlock" data-argument="h1" class="btn" value="H2" />' +
        '<input type="button" data-command="formatBlock" data-argument="h1" class="btn" value="H3" />' +
        '<input type="button" data-command="formatBlock" data-argument="h1" class="btn" value="H4" />' +
        '<input type="button" data-command="formatBlock" data-argument="p" class="btn" value="x">' +
        '&#160;&#160;' +
        '<input type="button" data-command="insertOrderedList" class="btn" value="OL">' +
        '<input type="button" data-command="insertUnorderedList" class="btn" value="UL">' +
        '&#160;&#160;' +
        '<input type="button" data-command="insertCode" class="btn" value="Insert code">' +
        '&#160;&#160;' +
     '</div>';

     var buttons = $$('.editor_buttons .btn');
     for(var i=0; i<buttons.length; i++)
         buttons[i].addEventListener("click", function(){
               var command = this.getAttribute('data-command');
               var argument = this.getAttribute('data-argument');

               if(command=="insertCode")
               {
                     document.execCommand("insertHTML", false, '<pre>\n\n</pre>');
                     return;
               }

               document.execCommand(this.getAttribute('data-command'), false, this.getAttribute('data-argument'));
           });

    $("#edit_btn").addEventListener("click", function(){
        $("#edit_btn").style.display = "none";
        $("#save_btn").style.display = "inline-block";
        $("#cancel_btn").style.display = "inline-block";

        $(".editor_buttons").style.display = "inline-block";


        var elements = $$("[data-editable=true][data-key]");
        for(var i=0; i<elements.length; i++)
        {
            var element = elements[i];
            element.contentEditable = true;
            element.addEventListener("keydown", editorKeyPress);
            element.addEventListener("focus", function(){
                if(this.getAttribute('data-panel')=="false")
                    $('.editor_buttons').style.display = "none";
                else
                    $('.editor_buttons').style.display = "inline-block";
            });

            element.style.outline = "none";

            element.focus();

            var timer_id = window.setInterval(function(){
                var elements = $$("[data-editable=true][data-key]");
                for(var i=0; i<elements.length; i++)
                {
                    var element = elements[i];
                    if(element.style.border=="none")
                    {
                        element.style.border = "1px solid red";
                        element.style.padding = "9px";
                    }
                    else
                    {
                        element.style.border = "none";
                        element.style.padding = "10px";
                    }
                }

            },600);

            element.setAttribute('timer-id', timer_id);
        }

    });
    $("#cancel_btn").addEventListener("click", function(){
        while(document.execCommand("undo")){}
        cancelEditor();
    });

    $("#save_btn").addEventListener("click", function(){
        cancelEditor();

        var elements = $$("[data-editable=true][data-key]");

        var data = new FormData();
        for(var i=0; i<elements.length; i++)
        {
            var element = elements[i];
            var key = element.getAttribute('data-key');
            var value = element.innerHTML;

            data.append(key, value);
        }

        var xhr = new XMLHttpRequest();

        xhr.onreadystatechange=function()
          {
          if (xhr.readyState==4 && xhr.status==200)
            {
//                window.location.reload();
            }
          }

        xhr.open("POST", "/api/post/", false);
//        xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
        xhr.send(data);
    });
}

function editorKeyPress(e) {
    var selection = window.getSelection();
    var eventNode = selection.getRangeAt(0).commonAncestorContainer.parentNode;

    // for <pre> tag
    if(eventNode.tagName=="PRE")
    {
        if(e.keyCode==18)
        {
            e.returnValue = false;
        }
        if(e.keyCode==13)
        {
            e.returnValue = false;
            document.execCommand("insertHTML", false, "\n");
        }
        if(e.keyCode==9) { // tab
            var selection = window.getSelection();
            e.returnValue = false;
            if(selection.toString().length>0)
            {
                document.execCommand("insertHTML", false, selection.toString().replace(/^([\s\S]*?)$/gm, "\t$1"));
            } else
                document.execCommand("insertHTML", false, "\t");
            return;
        }
        return
    }

    // for another tags
    if(e.keyCode==9) { // tab
        e.returnValue = false;
        document.execCommand("insertHTML", false, "&emsp;");
        return;
    }
}