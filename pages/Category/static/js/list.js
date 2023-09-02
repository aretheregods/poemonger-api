var deleteButtons = document.getElementsByClassName("delete");
for (var button of deleteButtons) {
    button.addEventListener("click", (e) => {
        e.preventDefault();
        console.log({ id: e.target.id });
    });
}
