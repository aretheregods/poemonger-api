import postData from "../../main/js/fetch.js";
import { handleSuccessResponse } from "../../main/js/postResponses.js";

postData(
    "/poetry",
    processFormData,
    handleSuccessResponse,
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
        categories.push(v.value);
    });

    fObj.categories = categories;
    fObj.poem = p.split("\n");
    fObj.releaseDate = new Date(...fObj.releaseDate.split("-"));
    fObj.sampleLength = Number(fObj.sampleLength);

    return JSON.stringify(fObj);
}
