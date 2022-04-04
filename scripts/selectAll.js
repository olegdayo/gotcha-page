function selectAll(source) {
    $(".ticks")
        .find("input[type=checkbox], select")
        .each(
            function () {
                this.checked = source.checked
            }
        )
}
