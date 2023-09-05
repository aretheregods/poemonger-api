export function handleSuccessResponse(r = {}) {
    var t = document.getElementById('success-message');
    var f = document.getElementById('form');
    var c = t.content.cloneNode(true);
    var a = c.querySelector('a');
    a.setAttribute('href', r.link);
    document.body.insertBefore(c, f);
}