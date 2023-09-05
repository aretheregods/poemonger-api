import postData from "../../main/js/fetch.js";

postData(
    "/poetry",
    processFormData,
    handlePostResponse,
	{ "Content-Type": "application/json" }
);

function processFormData(f) {
    var fObj = {};
    var c = document.querySelectorAll('#categories option:checked');
    var p = f.get("poem");
    var categories = [];
    f.forEach((v, k) => {
        fObj[k] = v;
    });
    c.forEach((v) => {
        categories.push({ id: v.value, name: v.id });
    });
    fObj.categories = categories;
    fObj.poem = p.split("\n");
    fObj.releaseDate = new Date(...fObj.releaseDate.split("-"));
    fObj.sampleLength = Number(fObj.sampleLength);
    return JSON.stringify(fObj);
}

function handlePostResponse(r = {}) {
    var t = document.getElementById('success-message');
    var f = document.getElementById('form');
    var c = t.content.cloneNode(true);
    var a = c.querySelector('a');
    a.setAttribute('href', r.link);
    document.body.insertBefore(c, f);
}
