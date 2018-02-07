function add_item() {
    $("#items_list").append('<li><p><input name="item[]" type="text" value="username"></p></li>');
}

function get_items() {
    var items = [];
    $("ol#items_list > li > p > input").each(function (index, element) {
        items.push($(element).val());
    });
    return items;
}

function mix(items) {
    var mixed_items = [];
    items.forEach(function (item, i) {
        mixed_items.splice(Math.floor(Math.random() * i), 0, item)
    });
    return mixed_items;
}

function hide(elem, duration) {
    elem.animate({opacity: 0.0}, duration);
    setTimeout(function(){elem.css("visibility", "hidden")}, duration);
}
function mix_queue() {
    var inputs = $("ol#items_list > li > p > input");
    mix(get_items()).forEach(function (value, i) {
        $(inputs[i]).val(value)
    });
    hide($("#btn-rand"), 1000);
}