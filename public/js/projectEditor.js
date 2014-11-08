$(document).ready(function () {
    var editor = new Editor({
        element: document.getElementById("project.Description")
    });
    editor.render();

    editor.codemirror.on("change", function(value) {
        value.getTextArea().value = (value.getValue());
    });
});