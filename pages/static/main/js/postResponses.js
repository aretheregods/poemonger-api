export function handleSuccessResponse(r = {}) {
    var t = document.getElementById('success-message');
    var f = document.getElementById('form');
    var c = t.content.cloneNode(true);
    var a = c.querySelector('a');

    a.setAttribute('href', r.link);
    f.parentNode.insertBefore(c, f);

    setTimeout(removeMessage, 7000);

    function removeMessage() {
        var m = document.getElementById('post-request-success');
    
        f.parentNode.removeChild(m);
        f.reset();
    }
}