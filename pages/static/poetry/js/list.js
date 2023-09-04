var deleteButtons = document.getElementsByClassName("delete");
for (var button of deleteButtons) {
    button.addEventListener("click", (e) => {
        e.preventDefault();
        var f = new FormData(undefined, e.target);
        f.append("id", e.target.id);
        fetch("/poetry", { method: "DELETE", body: f })
            .then((r) => window.location.reload())
            .catch(() =>
                console.error("There was an error deleting this entry.")
            );
    });
}