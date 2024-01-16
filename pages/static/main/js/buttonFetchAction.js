export default function buttonFetchAction({
    target = '',
    path = '',
    key = '',
    method = 'DELETE',
    successResponse = () => window.location.reload(),
    errorResponse = 'There was an error deleting this entry.',
    event = 'click',
    data = undefined,
}) {
    var deleteButtons = document.getElementsByClassName(target);
    for (var button of deleteButtons) {
        button.addEventListener(event, (e) => {
            e.preventDefault();
            var f = new FormData(data, e.target);

            f.append(key, e.target.id);

            fetch(`/${path}`, { method, body: f })
                .then((d) => successResponse(d))
                .catch(() =>
                    console.error(errorResponse)
                );
        });
    }
}