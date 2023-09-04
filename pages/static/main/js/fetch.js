export default function postData(path = '/', former = (form = FormData) => form, headers = {}) {
    var formElement = document.getElementById("form");
    if (!formElement) {
        throw "There is no form element";
    }
    var submitButton = document.getElementById("submit");
    if (!submitButton) {
        throw "There is no submit button";
    }
    formElement.addEventListener("submit", (e) => {
        e.preventDefault();
        var f = former(new FormData(formElement, submitButton))
        fetch(path, {
            method: "POST",
            body: f,
            headers
        })
            .then((r) => r.json())
            .then((r) => console.log(r))
            .catch((e) => console.error(e.json()));
    });
}