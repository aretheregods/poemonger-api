import postData from "../../main/js/fetch.js";

postData(
	"/poetry",
	(f) => {
		var fObj = {};
		var c = document.querySelectorAll('#categories option:checked');
        var p = f.get("poem");
        var categories = [];
		f.forEach((value, key) => {
			fObj[key] = value;
        });
        for (var v of c) {
            categories.push({ id: v.value, name: v.id });
        };
		fObj.categories = categories;
		fObj.poem = p.split("\n");
		fObj.release_date = new Date(...fObj.release_date.split("-"));
		console.log({ fObj });
		return JSON.stringify(fObj);
	},
	{ "Content-Type": "application/json" }
);
